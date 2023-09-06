package clerkauth

import (
	"context"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// key is used to store/retrieve a Claims value from a context.Context.
const claimKey ctxKey = 1

// =============================================================================

// SetClaims stores the claims in the context.
func SetClaims(ctx context.Context, claims ClerkClaims) context.Context {
	return context.WithValue(ctx, claimKey, claims)
}

// GetClaims returns the claims from the context.
func GetClaims(ctx context.Context) ClerkClaims {
	v, ok := ctx.Value(claimKey).(ClerkClaims)
	if !ok {
		return ClerkClaims{}
	}
	return v
}
