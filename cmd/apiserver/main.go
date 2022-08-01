package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"github.com/maxsnegir/rest-api-go/internal/app/db"
	"github.com/maxsnegir/rest-api-go/internal/app/server"
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

	pqStore := db.NewPqStore(config.PostgresDsn)
	if err := pqStore.Connect(); err != nil {
		log.Fatal(err)
	}

	logger.Infof("Connect to Postgres")
	defer func() {
		if err := pqStore.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	redisStore := db.NewRedisStore(config.RedisDsn)
	if err := redisStore.Connect(); err != nil {
		log.Fatal(err)
	}
	logger.Infof("Connect to Redis")
	defer func() {
		if err := redisStore.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	apiServer := server.NewServer(logger, config, pqStore.Connection, redisStore.Client)
	if err := apiServer.Start(); err != nil {
		log.Fatal(err)
	}

}
