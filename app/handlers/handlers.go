package handlers

import (
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/vim-diesel/new-service/app/handlers/v1/testgrp"
	"github.com/vim-diesel/new-service/foundation/web"
	"golang.org/x/exp/slog"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *slog.Logger
	DB       *sqlx.DB
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown)

	app.Handle(http.MethodGet, "/test", testgrp.Test)

	return app
}
