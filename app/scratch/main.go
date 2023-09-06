package main

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/golang-jwt/jwt/v4"
)

// GoogleClaims represents the authorization claims transmitted via a JWT.
type ClerkClaims struct {
	AuthorizedParty string `json:"azp,omitempty"`
	SessionID       string `json:"sid,omitempty"`
	jwt.StandardClaims
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {

	key, err := getClerkPubKey()
	if err != nil {
		return err
	}

	tokenString := `eyJhbGciOiJSUzI1NiIsImtpZCI6Imluc18yVWNVUk91NDFNamxUdDNzRlNOa0c0QW5vaTUiLCJ0eXAiOiJKV1QifQ.eyJhenAiOiJodHRwOi8vbG9jYWxob3N0OjUxNzMiLCJleHAiOjE2OTQwMjU4MjYsImlhdCI6MTY5NDAyNTc2NiwiaXNzIjoiaHR0cHM6Ly9zbWlsaW5nLXRvbWNhdC00OC5jbGVyay5hY2NvdW50cy5kZXYiLCJuYmYiOjE2OTQwMjU3NTYsInNpZCI6InNlc3NfMlYwSGNiank3bFllTHJhTzdrSFJYQ3JCODZnIiwic3ViIjoidXNlcl8yVjBIY1pCRnZuUlhwbG8xRWx6TkljZUdwWlUifQ.WsKSyAkTTN0BpbmgKSstmQ0OhEMKrtWt0Ubki3KUwr_X1TlXMxx_lqvPqEFkuby7-udp3Z2LCNLvrVsCJEjtxZKBIJCY0snDhLfhlDJagZx2nXkno8ennnc5WBgU1z5lEx6AT5vCaKqnNLb2jf7ci06cO1xHoqydq3tEaqYRdFby0bdYDwxwiAYckSaBK4UgqaSsXUObJI6upsLdttBHxQ0HtjU0GTf2trWF7DxKrM8q16BAtZVyJWMOUMvoFjAc-bH-iCMg-sVhGrZT6j8JUMKSN5pvOQjtaslVkRL9yYE-I2ChwzPk3WFLw1Fjc9k5EmgacvRzQyOl49YgORHIHQ`

	claimsStruct := ClerkClaims{}
	parser := jwt.NewParser(jwt.WithValidMethods([]string{jwt.SigningMethodRS256.Name}))

	token, err := parser.ParseWithClaims(
		tokenString,
		&claimsStruct,
		func(token *jwt.Token) (interface{}, error) {
			return key, nil
		},
	)
	if err != nil {
		return fmt.Errorf("parsing token: %w", err)
	}

	claims, ok := token.Claims.(*ClerkClaims)
	if !ok {
		return errors.New("Invalid Google JWT")
	}

	if claims.Issuer != "https://smiling-tomcat-48.clerk.accounts.dev" {
		return errors.New("iss is invalid")
	}

	if claims.AuthorizedParty != "http://localhost:5173" {
		return errors.New("iss is invalid")
	}

	return nil
}

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
