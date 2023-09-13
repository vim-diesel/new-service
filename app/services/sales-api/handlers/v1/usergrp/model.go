package usergrp

import (
	"fmt"
	"net/mail"
	"time"

	"github.com/vim-diesel/new-service/business/core/user"
	"github.com/vim-diesel/new-service/business/cview/user/summary"

	// "github.com/vim-diesel/new-service/business/cview/user/summary"
	"github.com/vim-diesel/new-service/foundation/validate"
)

// AppUser represents information about an individual user.
type AppUser struct {
	ID        string `json:"id"`
	FullName  string `json:"fullName"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Enabled   bool   `json:"enabled"`
	CreatedAt string `json:"createdAt"`
}

func toAppUser(usr user.User) AppUser {

	return AppUser{
		ID:        usr.ID,
		FullName:  usr.FullName,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email.Address,
		Enabled:   usr.Enabled,
		CreatedAt: usr.CreatedAt.Format(time.RFC3339),
	}
}

// =============================================================================

// AppNewUser contains information needed to create a new user.
type AppNewUser struct {
	FullName  string `json:"fullName" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
}

func toCoreNewUser(app AppNewUser) (user.NewUser, error) {

	addr, err := mail.ParseAddress(app.Email)
	if err != nil {
		return user.NewUser{}, fmt.Errorf("parsing email: %w", err)
	}

	usr := user.NewUser{
		FullName:  app.FullName,
		FirstName: app.FirstName,
		LastName:  app.LastName,
		Email:     *addr,
	}

	return usr, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewUser) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}
	return nil
}

// =============================================================================

// AppUpdateUser contains information needed to update a user.
type AppUpdateUser struct {
	FullName  *string `json:"fullName"`
	FirstName *string `json:"firstName"`
	LastName  *string `json:"lastName"`
	Email     *string `json:"email" validate:"omitempty,email"`
	Enabled   *bool   `json:"enabled"`
}

func toCoreUpdateUser(app AppUpdateUser) (user.UpdateUser, error) {

	var addr *mail.Address
	if app.Email != nil {
		var err error
		addr, err = mail.ParseAddress(*app.Email)
		if err != nil {
			return user.UpdateUser{}, fmt.Errorf("parsing email: %w", err)
		}
	}

	nu := user.UpdateUser{
		FullName:  app.FullName,
		FirstName: app.FirstName,
		LastName:  app.LastName,
		Email:     addr,
		Enabled:   app.Enabled,
	}

	return nu, nil
}

// Validate checks the data in the model is considered clean.
func (app AppUpdateUser) Validate() error {
	if err := validate.Check(app); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// =============================================================================

// AppSummary represents information about an individual user and their products.
type AppSummary struct {
	UserID     string `json:"userID"`
	UserName   string `json:"userName"`
	TotalCount int    `json:"totalCount"`
}

func toAppSummary(smm summary.Summary) AppSummary {
	return AppSummary{
		UserID:     smm.UserID.String(),
		UserName:   smm.UserName,
		TotalCount: smm.TotalCount,
	}
}
