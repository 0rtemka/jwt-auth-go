package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"test"
	"test/pkg/handler"
	"test/pkg/repository"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "configPath", "configs/config.toml", "path to config toml file")
}

func main() {
	flag.Parse()
	config, err := test.NewConfig(configPath)
	if err != nil {
		log.Fatalf("error initializing config: %s", err.Error())
	}

	db := repository.NewMongoDB(config.Mongo.URI, config.Mongo.Database)
	print(db)

	router := handler.NewHandler()
	server := new(test.Server)
	log.Infof("starting server on port: %s", config.Port)

	if err := server.Run(config.Port, router.InitRoutes()); err != nil {
		log.Fatalf("error running server: %s", err.Error())
	}
}
