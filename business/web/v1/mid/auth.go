package mid

import (
	"context"
	"net/http"

	"github.com/vim-diesel/new-service/business/web/clerkauth"
	"github.com/vim-diesel/new-service/foundation/web"
)

func Authenticate(a *clerkauth.ClerkAuth) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			claims, err := a.ValidateClerkJWT(ctx, r.Header.Get("authorization"))
			if err != nil {
				return clerkauth.NewAuthError("authenticate: failed: %s", err)
			}

			ctx = clerkauth.SetClaims(ctx, claims)
			ctx = clerkauth.SetUserID(ctx, claims.Subject)
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}

// Authorize validates that an authenticated user has claims
// func Authorize(a *clerkauth.ClerkAuth) web.Middleware {
// 	m := func(handler web.Handler) web.Handler {
// 		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 			claims := clerkauth.GetClaims(ctx)
// 			if claims.Subject == "" {
// 				return clerkauth.NewAuthError("authorize: you are not authorized for that action, no claims")
// 			}

// 			ctx = clerkauth.SetUserID(ctx, claims.Subject)
// 			return handler(ctx, w, r)
// 		}

// 		return h
// 	}

// 	return m
// }
