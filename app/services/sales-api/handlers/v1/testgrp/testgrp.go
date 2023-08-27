package testgrp

import (
	"context"
	"net/http"

	"github.com/vim-diesel/new-service/foundation/web"
)

// Test is our example route.
func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	// if n := rand.Intn(100); n%2 == 0 {
	// 	return v1.NewRequestError(errors.New("TRUSTED ERROR"), http.StatusBadRequest)
	// }

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

	status := struct {
		Status  string
		Message string
	}{
		Status:  "OK",
		Message: "You are authorized to view this page",
	}

	return web.Respond(ctx, w, status, http.StatusOK)
}
