package googauth

import (
	"context"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// key is used to store/retrieve a Claims value from a context.Context.
const claimKey ctxKey = 1

// =============================================================================

// SetClaims stores the claims in the context.
func SetClaims(ctx context.Context, claims GoogleClaims) context.Context {
	return context.WithValue(ctx, claimKey, claims)
}

// GetClaims returns the claims from the context.
func GetClaims(ctx context.Context) GoogleClaims {
	v, ok := ctx.Value(claimKey).(GoogleClaims)
	if !ok {
		return GoogleClaims{}
	}
	return v
}
