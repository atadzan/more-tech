package app

import (
	"context"
	"fmt"
	"github.com/atadzan/more-tech/app/server"
	"log"
	"os"
)

func Init(ctx context.Context) error {

	dbConfig := repository.Config{
		Host:     os.Getenv(constants.DB_HOST),
		Username: os.Getenv(constants.DB_USER),
		Password: os.Getenv(constants.DB_PASSWORD),
		DBName:   os.Getenv(constants.DB_NAME),
		SSLMode:  os.Getenv(constants.DB_SSL),
	}

	dbPool, err := repository.NewPostgresDB(ctx, dbConfig)

	if err != nil {
		debug.Trace()
		return fmt.Errorf("failed to open database: %v", err)
	}

	defer dbPool.Close()

	repo := repository.NewRepository(dbPool)
	service := services.NewService(repo)
	handlers := v1.NewHandler(service)
	srv := new(server.Server)

	if err = srv.Run(os.Getenv(constants.HTTP_PORT), handlers.InitRoutes()); err != nil {
		debug.Trace()
		log.Fatalf("error http server: %s", err.Error())
	}

	return nil
}
