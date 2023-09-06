package clerkauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"time"

	"github.com/clerkinc/clerk-sdk-go/clerk"
	"github.com/golang-jwt/jwt/v4"
)

// ErrForbidden is returned when a auth issue is identified.
var ErrForbidden = errors.New("attempted action is not allowed")

type ClerkClaims struct {
	jwt.StandardClaims
}

// Config represents information required to initialize auth.
type Config struct {
	Log    *slog.Logger
	Client *clerk.Client
}

// Auth is used to authenticate clients.
type ClerkAuth struct {
	log    *slog.Logger
	client *clerk.Client
}

// New creates an ClerkAuth to support authentication.
func New(cfg Config) (*ClerkAuth, error) {
	a := ClerkAuth{
		log:    cfg.Log,
		client: cfg.Client,
	}

	return &a, nil
}

// ValidateGoogleJWT -
func (a *ClerkAuth) ValidateClerkJWT(tokenString string) (ClerkClaims, error) {

	parts := strings.Split(tokenString, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ClerkClaims{}, errors.New("expected authorization header format: Bearer <token>")
	}

	claimsStruct := ClerkClaims{}

	token, err := a.parser.ParseWithClaims(
		parts[1],
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
		return ClerkClaims{}, err
	}

	claims, ok := token.Claims.(*ClerkClaims)
	if !ok {
		return ClerkClaims{}, errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "accounts.google.com" && claims.Issuer != "https://accounts.google.com" {
		return ClerkClaims{}, errors.New("iss is invalid")
	}

	if claims.Audience != a.audience {
		return ClerkClaims{}, errors.New("aud is invalid")
	}

	if claims.ExpiresAt < time.Now().UTC().Unix() {
		return ClerkClaims{}, errors.New("JWT is expired")
	}

	return *claims, nil
}

// =============================================================================

func getGooglePublicKey(keyID string) (string, error) {
	resp, err := http.Get("https://smiling-tomcat-48.clerk.accounts.dev/.well-known/jwks.json")
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

// func getClerkPubKey() (*rsa.PublicKey, error) {

// 	ctx := context.Background()
// 	// Fetch the JWK from the URL
// 	set, err := jwk.Fetch(ctx, "https://smiling-tomcat-48.clerk.accounts.dev/.well-known/jwks.json")
// 	if err != nil {
// 		return nil, err
// 	}

// 	fmt.Println(set)

// 	var rsa *rsa.PublicKey

// 	for it := set.Iterate(context.Background()); it.Next(context.Background()); {
// 		pair := it.Pair()
// 		key := pair.Value.(jwk.Key)

// 		var rawkey interface{} // This is the raw key, like *rsa.PrivateKey or *ecdsa.PrivateKey
// 		if err := key.Raw(&rawkey); err != nil {
// 			log.Printf("failed to create public key: %s", err)
// 			return nil, err
// 		}

// 		// We know this is an RSA Key so...
// 		rsa, ok := &rawkey.(*rsa.PublicKey)
// 		if !ok {
// 			panic(fmt.Sprintf("expected ras key, got %T", rawkey))
// 		}
// 	}
// 	return rsa, nil
// }
