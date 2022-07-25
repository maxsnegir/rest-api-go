package server

import (
	"github.com/maxsnegir/rest-api-go/internal/app/store"
	"github.com/sirupsen/logrus"
)

type apiServer struct {
	Logger *logrus.Logger
	Config *Config
	Store  store.Store
}

func (s *apiServer) Start() {
	s.Logger.Info("Server is running")
}

func NewServer(logger *logrus.Logger, config *Config, store store.Store) *apiServer {
	return &apiServer{
		Logger: logger,
		Config: config,
		Store:  store,
	}
}
