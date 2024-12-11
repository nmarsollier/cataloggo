package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nmarsollier/cataloggo/tools/env"
	"github.com/nmarsollier/cataloggo/tools/log"
)

var (
	instance *pgxpool.Pool
)

func GetPostgresClient(deps ...interface{}) (*pgxpool.Pool, error) {
	if instance == nil {
		config, err := pgxpool.ParseConfig(env.Get().PostgresURL)
		if err != nil {
			log.Get(deps...).Error(err)
			return nil, err
		}

		instance, err = pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Get(deps...).Error(err)
			return nil, err
		}

		_, err = instance.Exec(context.Background(), "SET search_path TO cataloggo")
		if err != nil {
			log.Get(deps...).Error(err)
		}

		log.Get(deps...).Info("Postgres Connected")
	}

	return instance, nil
}
