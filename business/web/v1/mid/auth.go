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
			return handler(ctx, w, r)
		}
		return h
	}
	return m
}

// Authenticate validates a JWT from the `Authorization` header.
// func Authenticate(a *googauth.GoogAuth) web.Middleware {
// 	m := func(handler web.Handler) web.Handler {
// 		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 			claims, err := a.ValidateGoogleJWT(r.Header.Get("authorization"))
// 			if err != nil {
// 				return googauth.NewAuthError("authenticate: failed: %s", err)
// 			}

// 			ctx = googauth.SetClaims(ctx, claims)

// 			return handler(ctx, w, r)
// 		}

// 		return h
// 	}

// 	return m
// }

// Authorize validates that an authenticated user has at least one role from a
// specified list. This method constructs the actual function that is used.
// func Authorize(a *googauth.GoogAuth, rule string) web.Middleware {
// 	m := func(handler web.Handler) web.Handler {
// 		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
// 			claims := auth.GetClaims(ctx)
// 			if claims.Subject == "" {
// 				return auth.NewAuthError("authorize: you are not authorized for that action, no claims")
// 			}

// 			if err := a.Authorize(ctx, claims, rule); err != nil {
// 				return auth.NewAuthError("authorize: you are not authorized for that action, claims[%v] rule[%v]: %s", claims.Roles, rule, err)
// 			}

// 			return handler(ctx, w, r)
// 		}

// 		return h
// 	}

// 	return m
// }
