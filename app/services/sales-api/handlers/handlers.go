package handlers

import (
	"net/http"
	"os"

	"log/slog"

	"github.com/jmoiron/sqlx"
	v1 "github.com/vim-diesel/new-service/app/services/sales-api/handlers/v1"
	"github.com/vim-diesel/new-service/business/web/clerkauth"
	"github.com/vim-diesel/new-service/business/web/v1/mid"
	"github.com/vim-diesel/new-service/foundation/web"
)

// Options represent optional parameters.
type Options struct {
	corsOrigin string
}

// WithCORS provides configuration options for CORS.
func WithCORS(origin string) func(opts *Options) {
	return func(opts *Options) {
		opts.corsOrigin = origin
	}
}

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Build     string
	Shutdown  chan os.Signal
	Log       *slog.Logger
	DB        *sqlx.DB
	ClerkAuth *clerkauth.ClerkAuth
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig, options ...func(opts *Options)) http.Handler {
	var opts Options
	for _, option := range options {
		option(&opts)
	}

	app := web.NewApp(
		cfg.Shutdown,
		mid.Logger(cfg.Log),
		mid.Errors(cfg.Log),
		mid.Panics(),
	)

	if opts.corsOrigin != "" {
		app.EnableCORS(mid.Cors(opts.corsOrigin))
	}

	v1.Routes(app, v1.Config{
		Build:     cfg.Build,
		Log:       cfg.Log,
		ClerkAuth: cfg.ClerkAuth,
		DB:        cfg.DB,
	})

	return app
}
