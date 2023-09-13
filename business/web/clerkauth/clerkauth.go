package clerkauth

import (
	"context"
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// ErrForbidden is returned when a auth issue is identified.
var ErrForbidden = errors.New("attempted action is not allowed")

type ClerkClaims struct {
	AuthorizedParty string `json:"azp,omitempty"`
	SessionID       string `json:"sid,omitempty"`
	jwt.StandardClaims
}

// Config represents information required to initialize auth.
type Config struct {
	Log *slog.Logger
}

// Auth is used to authenticate clients.
type ClerkAuth struct {
	log    *slog.Logger
	parser *jwt.Parser
}

// New creates an ClerkAuth to support authentication.
func New(cfg Config) (*ClerkAuth, error) {
	a := ClerkAuth{
		log:    cfg.Log,
		parser: jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name})),
	}

	return &a, nil
}

// ValidateGoogleJWT -
func (a *ClerkAuth) ValidateClerkJWT(ctx context.Context, tokenString string) (ClerkClaims, error) {

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ClerkClaims{}, errors.New("expected authorization header format: Bearer <token>")
	}

	claimsStruct := ClerkClaims{}

	token, err := a.parser.ParseWithClaims(
		parts[1],
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			key, err := getClerkPubKey()
			if err != nil {
				return nil, fmt.Errorf("getting public key: %w", err)
			}
			return key, nil
		},
	)
	if err != nil {
		return ClerkClaims{}, fmt.Errorf("parsing JWT: %w", err)
	}

	claims, ok := token.Claims.(*ClerkClaims)
	if !ok {
		return ClerkClaims{}, errors.New("Invalid Clerk JWT")
	}

	if claims.Issuer != "https://smiling-tomcat-48.clerk.accounts.dev" {
		return ClerkClaims{}, errors.New("iss is invalid")
	}

	if claims.AuthorizedParty != "http://localhost:5173" && claims.AuthorizedParty != "https://new-service.vercel.app" {
		return ClerkClaims{}, errors.New("azp is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return ClerkClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}

// =============================================================================

func getClerkPubKey() (*rsa.PublicKey, error) {
	file, err := os.Open("key.pem")

	if err != nil {
		return nil, fmt.Errorf("reading opening file: %w", err)
	}

	defer file.Close()

	pemData, err := io.ReadAll(io.LimitReader(file, 1024*1024))
	if err != nil {
		return nil, fmt.Errorf("reading auth public key: %w", err)
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(pemData)
	if err != nil {
		return nil, fmt.Errorf("parsing auth public key: %w", err)
	}

	return publicKey, nil
}
