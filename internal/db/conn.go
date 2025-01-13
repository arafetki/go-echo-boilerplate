package db

import (
	"context"
	"errors"
	"time"

	"github.com/arafetki/go-echo-boilerplate/assets"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
)

type DB struct {
	*pgxpool.Pool
}

func Pool(dsn string, automigrate bool) (*DB, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	config, err := pgxpool.ParseConfig("postgresql://" + dsn)
	if err != nil {
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	if automigrate {
		iofsDriver, err := iofs.New(assets.Migrations, "migrations")
		if err != nil {
			return nil, err
		}
		migrator, err := migrate.NewWithSourceInstance("iofs", iofsDriver, "postgresql://"+dsn)
		if err != nil {
			return nil, err
		}
		err = migrator.Up()
		if err != nil {
			switch {
			case errors.Is(err, migrate.ErrNoChange):
				break
			default:
				return nil, err
			}
		}
	}

	return &DB{pool}, nil
}
