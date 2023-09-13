package userdb

import (
	"net/mail"
	"time"

	"github.com/vim-diesel/new-service/business/core/user"
)

// dbUser represent the structure we need for moving data
// between the app and the database.
type dbUser struct {
	ID        string    `db:"user_id"`
	FullName  string    `db:"full_name"`
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	Email     string    `db:"email"`
	Enabled   bool      `db:"enabled"`
	CreatedAt time.Time `db:"created_at"`
}

func toDBUser(usr user.User) dbUser {

	return dbUser{
		ID:        usr.ID,
		FullName:  usr.FullName,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email.Address,
		Enabled:   usr.Enabled,
		CreatedAt: usr.CreatedAt.UTC(),
	}
}

func toCoreUser(dbUsr dbUser) user.User {
	addr := mail.Address{
		Address: dbUsr.Email,
	}

	usr := user.User{
		ID:        dbUsr.ID,
		FullName:  dbUsr.FullName,
		FirstName: dbUsr.FirstName,
		LastName:  dbUsr.LastName,
		Email:     addr,
		Enabled:   dbUsr.Enabled,
		CreatedAt: dbUsr.CreatedAt.In(time.Local),
	}

	return usr
}

func toCoreUserSlice(dbUsers []dbUser) []user.User {
	usrs := make([]user.User, len(dbUsers))
	for i, dbUsr := range dbUsers {
		usrs[i] = toCoreUser(dbUsr)
	}
	return usrs
}
