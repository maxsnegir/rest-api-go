package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/maxsnegir/rest-api-go/internal/app/server"
	"github.com/maxsnegir/rest-api-go/internal/app/store"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		log.Fatal(err)
	}
	logger, err := server.NewLogger(config.LogLevel)
	if err != nil {
		log.Fatal(err)
	}

	pqStore := store.NewStore("postgres", config.DatabaseUrl)
	if err := pqStore.Connect(); err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := pqStore.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	apiServer := server.NewServer(logger, config, pqStore)
	apiServer.Start()

}
