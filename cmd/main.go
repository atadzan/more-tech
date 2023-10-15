package main

import (
	"context"
	"flag"
	"github.com/atadzan/more-tech/app"
	"log"
)

var configPath *string

func init() {
	// app configuration path flag
	configPath = flag.String("config", "./config/config.yaml", "Default config path")
	flag.Parse()
}

func main() {
	if err := app.Init(context.Background(), *configPath); err != nil {
		log.Println(err)
		return
	}
}
