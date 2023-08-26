package main

import (
	"context"
	"errors"
	"expvar"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"log/slog"

	"github.com/ardanlabs/conf/v3"

	"github.com/joho/godotenv"
	"github.com/vim-diesel/new-service/app/services/sales-api/handlers"
	database "github.com/vim-diesel/new-service/business/sys/database/pgx"
	"github.com/vim-diesel/new-service/business/web/googauth"
	"github.com/vim-diesel/new-service/business/web/v1/debug"
)

// TODO:
// - database.Open() is not returning an error when we fail to get a DSN from
//   .env. This is the sqlx package so there probably isn't anything to do about
//   this, just something to be aware of.

var build = "develop"

func main() {
	log := slog.Default()
	ctx := context.Background()

	if err := run(ctx, log); err != nil {
		log.ErrorContext(ctx, "startup", "msg", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, log *slog.Logger) error {
	// -------------------------------------------------------------------------
	// GOMAXPROCS

	log.InfoContext(ctx, "startup", "GOMAXPROCS", runtime.GOMAXPROCS(0))

	// -------------------------------------------------------------------------
	// Configuration

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("godotenv: %w", err)
	}

	dsn := os.Getenv("DSN")
	audience := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")

	cfg := struct {
		conf.Version
		Web struct {
			ReadTimeout     time.Duration `conf:"default:10s"`
			WriteTimeout    time.Duration `conf:"default:15s"`
			IdleTimeout     time.Duration `conf:"default:120s"`
			ShutdownTimeout time.Duration `conf:"default:20s"`
			APIHost         string        `conf:"default:0.0.0.0:3000"`
			DebugHost       string        `conf:"default:0.0.0.0:4000"`
		}
	}{
		Version: conf.Version{
			Build: build,
			Desc:  "VIM DIESEL",
		},
	}

	const prefix = "SALES"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}
		return fmt.Errorf("parsing config: %w", err)
	}

	// -------------------------------------------------------------------------
	// App Starting

	log.InfoContext(ctx, "starting service", "version", build)
	defer log.InfoContext(ctx, "shutdown complete")

	expvar.NewString("build").Set(build)

	// -------------------------------------------------------------------------
	// Database Support

	log.InfoContext(ctx, "startup", "status", "initializing neon.tech database support")

	db, err := database.Open(dsn, 2, 112)
	if err != nil {
		return fmt.Errorf("failed to initialize a connection to neon.tech database: %w", err)
	}

	// If we need to, we can increase the deadline here. Is quite slow, because
	// this is a serverless solution, so cold starts are a thing.
	pingDeadline := time.Duration(15 * time.Second)

	if err := database.StatusCheck(ctx, db, pingDeadline); err != nil {
		return fmt.Errorf("database statuscheck: %w", err)
	}

	defer func() {
		log.InfoContext(ctx, "shutdown", "status", "stopping database support")
		db.Close()
	}()

	// -------------------------------------------------------------------------
	// Initialize authentication support

	log.InfoContext(ctx, "startup", "status", "initializing authentication support")

	googAuthCfg := googauth.Config{
		Log:      log,
		Issuer:   "https://accounts.google.com",
		Audience: audience,
	}

	googAuth, err := googauth.New(googAuthCfg)
	if err != nil {
		return fmt.Errorf("constructing auth: %w", err)
	}

	// -------------------------------------------------------------------------
	// Start Debug Service

	log.InfoContext(ctx, "startup", "status", "debug v1 router started", "host", cfg.Web.DebugHost)

	go func() {
		if err := http.ListenAndServe(cfg.Web.DebugHost, debug.StandardLibraryMux()); err != nil {
			log.ErrorContext(ctx, "shutdown", "status", "debug v1 router closed", "host", cfg.Web.DebugHost, "ERROR", err)
		}
	}()

	// -------------------------------------------------------------------------
	// Start API Service

	log.InfoContext(ctx, "startup", "status", "initializing V1 API support")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	cfgMux := handlers.APIMuxConfig{
		Build:    build,
		Shutdown: shutdown,
		Log:      log,
		DB:       db,
		GoogAuth: googAuth,
	}
	apiMux := handlers.APIMux(cfgMux, handlers.WithCORS("*"))

	api := http.Server{
		Addr:         cfg.Web.APIHost,
		Handler:      apiMux,
		ReadTimeout:  cfg.Web.ReadTimeout,
		WriteTimeout: cfg.Web.WriteTimeout,
		IdleTimeout:  cfg.Web.IdleTimeout,
		ErrorLog:     slog.NewLogLogger(log.Handler(), slog.LevelError),
	}

	serverErrors := make(chan error, 1)

	go func() {
		log.InfoContext(ctx, "startup", "status", "api router started", "host", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	// -------------------------------------------------------------------------
	// Shutdown

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		log.InfoContext(ctx, "shutdown", "status", "shutdown started", "signal", sig)
		defer log.InfoContext(ctx, "shutdown", "status", "shutdown complete", "signal", sig)

		ctx, cancel := context.WithTimeout(context.Background(), cfg.Web.ShutdownTimeout)
		defer cancel()

		if err := api.Shutdown(ctx); err != nil {
			api.Close()
			return fmt.Errorf("could not stop server gracefully: %w", err)
		}
	}

	return nil
}
