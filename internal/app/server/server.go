package server

import (
	"github.com/gorilla/mux"
	"github.com/maxsnegir/rest-api-go/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type apiServer struct {
	Logger *logrus.Logger
	Config *Config
	Store  store.Store
	Router *mux.Router
}

func (s *apiServer) Start() error {
	s.Logger.Infof("Starting server on http://localhost%s", s.Config.Port)
	err := http.ListenAndServe(s.Config.Port, s.Router)
	if err != nil {
		return err
	}
	return nil
}

func NewServer(logger *logrus.Logger, config *Config, store store.Store) *apiServer {
	return &apiServer{
		Logger: logger,
		Config: config,
		Store:  store,
		Router: mux.NewRouter(),
	}
}
