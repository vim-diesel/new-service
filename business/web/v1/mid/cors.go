package mid

import (
	"context"
	"net/http"

	"github.com/vim-diesel/new-service/foundation/web"
)

// Cors sets the response headers needed for Cross-Origin Resource Sharing
func Cors(origin string) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			w.Header().Set("Access-Control-Max-Age", "86400")
			// w.Header().Set("Access-Control-Allow-Credentials", "true")

			return handler(ctx, w, r)
		}

		return h
	}

	return m
}
