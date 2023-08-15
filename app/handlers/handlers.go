package handlers

import (
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/vim-diesel/new-service/app/handlers/v1/testgrp"
	"github.com/vim-diesel/new-service/app/handlers/v1/usergrp"
	"github.com/vim-diesel/new-service/business/core/user"
	"github.com/vim-diesel/new-service/business/core/user/stores/userdb"
	"github.com/vim-diesel/new-service/business/web/v1/mid"
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
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Panics())

	app.Handle(http.MethodGet, "/test", testgrp.Test)

	// =========================================================================

	usrCore := user.NewCore(userdb.NewStore(cfg.Log, cfg.DB))

	ugh := usergrp.New(usrCore)

	app.Handle(http.MethodGet, "/users", ugh.Query)

	return app
}
