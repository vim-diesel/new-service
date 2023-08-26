package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	v1 "github.com/vim-diesel/new-service/business/web/v1"
	"github.com/vim-diesel/new-service/foundation/web"
)

// Test is our example route.
func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if n := rand.Intn(100); n%2 == 0 {
		return v1.NewRequestError(errors.New("TRUSTED ERROR"), http.StatusBadRequest)
	}

	// Validate the data
	// Call into the business layer

	status := struct {
		Status string
	}{
		Status: "OK",
	}

	return web.Respond(ctx, w, status, http.StatusOK)

}

func TestingAuth(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	status := struct {
		Status  string
		Message string
	}{
		Status:  "OK",
		Message: "You are authorized to view this page",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
