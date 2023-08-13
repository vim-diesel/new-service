package dbmigrate

import "golang.org/x/exp/slog"

type dbSeeder struct {
	log *slog.Logger
}

const (
	// There's a transaction limit for PlanetScale, so we're limited on the amount of seed data we can create
	usersToSeed = 10
)
