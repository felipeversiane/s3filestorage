package database

import (
	"context"
	"fmt"
	"time"

	"github.com/felipeversiane/s3filestorage/internal/infra/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Connection *pgxpool.Pool

func Connect(ctx context.Context) error {
	dsn := getConnectionString()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return err
	}

	Connection, err = pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return err
	}
	return nil
}

func Close() {
	if Connection == nil {
		return
	}
	Connection.Close()
}

func getConnectionString() string {
	user := config.Conf.Database.User
	password := config.Conf.Database.Password
	dbname := config.Conf.Database.Name
	dbport := config.Conf.Database.Port
	dbhost := config.Conf.Database.Host

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable", user, password, dbname, dbport, dbhost)

	return dsn
}
