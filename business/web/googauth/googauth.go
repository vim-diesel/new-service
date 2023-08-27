// Package auth provides authentication and authorization support.
// Authentication: You are who you say you are.
// Authorization:  You have permission to do what you are requesting to do.
package googauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// ErrForbidden is returned when a auth issue is identified.
var ErrForbidden = errors.New("attempted action is not allowed")

// GoogleClaims represents the authorization claims transmitted via a JWT.
type GoogleClaims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	FirstName     string `json:"given_name"`
	LastName      string `json:"family_name"`
	jwt.StandardClaims
}

// Config represents information required to initialize auth.
type Config struct {
	Log      *slog.Logger
	Issuer   string
	Audience string
}

// Auth is used to authenticate clients. It can generate a token for a
// set of user claims and recreate the claims by parsing the token.
type GoogAuth struct {
	log      *slog.Logger
	method   jwt.SigningMethod
	parser   *jwt.Parser
	issuer   string
	audience string
}

// New creates an GoogAuth to support authentication/authorization.
func New(cfg Config) (*GoogAuth, error) {
	a := GoogAuth{
		log:      cfg.Log,
		method:   jwt.GetSigningMethod(jwt.SigningMethodRS256.Name),
		parser:   jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name})),
		issuer:   cfg.Issuer,
		audience: cfg.Audience,
	}

	return &a, nil
}

// ValidateGoogleJWT -
func (a *GoogAuth) ValidateGoogleJWT(tokenString string) (GoogleClaims, error) {
	claimsStruct := GoogleClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			pem, err := getGooglePublicKey(fmt.Sprintf("%s", token.Header["kid"]))
			if err != nil {
				return nil, err
			}
			key, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pem))
			if err != nil {
				return nil, err
			}
			return key, nil
		},
	)
	if err != nil {
		return GoogleClaims{}, err
	}

	claims, ok := token.Claims.(*GoogleClaims)
	if !ok {
		return GoogleClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != a.issuer {
		return GoogleClaims{}, errors.New("iss is invalid")
	}

	if claims.Audience != a.audience {
		return GoogleClaims{}, errors.New("aud is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return GoogleClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}

// Authorize attempts to authorize the user with the provided input roles, if
// none of the input roles are within the user's claims, we return an error
// otherwise the user is authorized.
// func (a *Auth) Authorize(ctx context.Context, claims Claims, rule string) error {
// 	input := map[string]any{
// 		"Roles":   claims.Roles,
// 		"Subject": claims.Subject,
// 		"UserID":  claims.Subject,
// 	}

// 	if err := a.opaPolicyEvaluation(ctx, opaAuthorization, rule, input); err != nil {
// 		return fmt.Errorf("rego evaluation failed : %w", err)
// 	}

// 	return nil
// }

// =============================================================================
func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v1/certs")
	if err != nil {
		return "", err
	}
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	myResp := map[string]string{}
	err = json.Unmarshal(dat, &myResp)
	if err != nil {
		return "", err
	}
	key, ok := myResp[keyID]
	if !ok {
		return "", errors.New("key not found")
	}
	return key, nil
}
