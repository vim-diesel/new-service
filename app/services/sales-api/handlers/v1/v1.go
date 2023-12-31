// Package v1 contains the full set of handler functions and routes
// supported by the v1 web api.
package v1

import (
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/vim-diesel/new-service/app/services/sales-api/handlers/v1/testgrp"
	"github.com/vim-diesel/new-service/app/services/sales-api/handlers/v1/usergrp"
	"github.com/vim-diesel/new-service/business/core/user"
	"github.com/vim-diesel/new-service/business/core/user/stores/userdb"
	"github.com/vim-diesel/new-service/business/web/clerkauth"
	"github.com/vim-diesel/new-service/business/web/v1/mid"
	"github.com/vim-diesel/new-service/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build     string
	Log       *slog.Logger
	ClerkAuth *clerkauth.ClerkAuth
	DB        *sqlx.DB
}

// Routes binds all the version 1 routes.
func Routes(app *web.App, cfg Config) {
	authen := mid.Authenticate(cfg.ClerkAuth)

	app.Handle(http.MethodGet, "/test", testgrp.Test)
	app.Handle(http.MethodGet, "/test/auth", testgrp.TestingAuth, authen)

	// =========================================================================

	usrCore := user.NewCore(userdb.NewStore(cfg.Log, cfg.DB))

	ugh := usergrp.New(usrCore)

	app.Handle(http.MethodGet, "/users", ugh.Query)
	app.Handle(http.MethodPost, "/users/create", ugh.Create, authen)
	app.Handle(http.MethodDelete, "/users/delete", ugh.Delete, authen)
}
