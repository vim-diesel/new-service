package handlers

import (
	"net/http"
	"os"

	"log/slog"

	"github.com/jmoiron/sqlx"
	"github.com/vim-diesel/new-service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/vim-diesel/new-service/app/services/sales-api/handlers/v1/usergrp"
	"github.com/vim-diesel/new-service/business/core/user"
	"github.com/vim-diesel/new-service/business/core/user/stores/userdb"
	"github.com/vim-diesel/new-service/business/web/auth"
	"github.com/vim-diesel/new-service/business/web/v1/mid"
	"github.com/vim-diesel/new-service/foundation/web"
)

// APIMuxConfig contains all the mandatory systems required by handlers.
type APIMuxConfig struct {
	Shutdown chan os.Signal
	Log      *slog.Logger
	DB       *sqlx.DB
	Auth     *auth.Auth
}

// APIMux constructs a http.Handler with all application routes defined.
func APIMux(cfg APIMuxConfig) *web.App {
	app := web.NewApp(cfg.Shutdown, mid.Logger(cfg.Log), mid.Errors(cfg.Log), mid.Panics())

	app.Handle(http.MethodGet, "/test", testgrp.Test)
	app.Handle(http.MethodGet, "/login", testgrp.TestingAuth)

	// =========================================================================

	usrCore := user.NewCore(userdb.NewStore(cfg.Log, cfg.DB))

	ugh := usergrp.New(usrCore)

	app.Handle(http.MethodGet, "/users", ugh.Query)

	return app
}
