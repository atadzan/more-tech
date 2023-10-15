package app

import (
	"context"
	"github.com/atadzan/more-tech/app/server"
	"github.com/atadzan/more-tech/config"
	"github.com/atadzan/more-tech/pkg/handler"
	"github.com/atadzan/more-tech/pkg/repository"
	"github.com/atadzan/more-tech/pkg/services"
	"github.com/atadzan/more-tech/third_party/database"
	"log"
)

func Init(ctx context.Context, cfgPath string) error {
	appCfg, err := config.LoadConfig(cfgPath)
	if err != nil {
		log.Println(err)
		return err
	}

	dbPoolConn, err := database.NewPostgresConn(ctx, appCfg.Postgres)
	if err != nil {
		log.Println(err)
		return err
	}

	repo := repository.NewRepository(dbPoolConn)
	service := services.NewService(repo)
	handlers := handler.NewHandler(service)
	srv := new(server.Server)

	if err = srv.Run(":"+appCfg.HTTP.Port, handlers.InitRoutes()); err != nil {
		log.Fatalf("error http server: %s", err.Error())
	}

	return nil
}
