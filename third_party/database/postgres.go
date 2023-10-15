package database

import (
	"context"
	"github.com/atadzan/more-tech/config"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/url"
)

func NewPostgresConn(ctx context.Context, cfg config.Postgres) (*pgxpool.Pool, error) {
	connUrl := "postgres://" + cfg.Username + ":" + url.QueryEscape(cfg.Password) + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName + "?sslmode=" + cfg.SSLMode
	dbConn, err := pgxpool.Connect(ctx, connUrl)
	if err != nil {
		return nil, err
	}
	err = dbConn.Ping(ctx)
	if err != nil {
		return nil, err
	}
	return dbConn, nil

}
