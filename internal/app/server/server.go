package server

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/maxsnegir/rest-api-go/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type apiServer struct {
	logger       *logrus.Logger
	Config       *Config
	PgConnection *sql.DB
	UserStore    store.UserStore
	router       *mux.Router
}

func (s *apiServer) beforeStart() {
	s.UserStore = store.NewUserStore(s.PgConnection)

}

func (s *apiServer) configureRouter() {
	s.router.Use(loggingMiddleware)
	s.router.HandleFunc("/user/create", s.CreateUser()).Methods(http.MethodPost)
}

func (s *apiServer) Start() error {
	s.beforeStart()
	s.configureRouter()
	s.logger.Infof("Starting server on http://localhost%s", s.Config.Port)
	err := http.ListenAndServe(s.Config.Port, s.router)
	if err != nil {
		return err
	}
	return nil
}

func NewServer(logger *logrus.Logger, config *Config, pqStore *sql.DB) *apiServer {
	return &apiServer{
		logger:       logger,
		Config:       config,
		PgConnection: pqStore,
		router:       mux.NewRouter(),
	}
}
