package server

import (
	"database/sql"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"github.com/maxsnegir/rest-api-go/internal/app/services"
	"github.com/sirupsen/logrus"
	"net/http"
)

type apiServer struct {
	logger *logrus.Logger
	Config *Config
	router *mux.Router
	// Db connections
	PgConnection *sql.DB
	RedisClient  *redis.Client
	// Services for work with API
	UserStore    services.UserStore
	TokenService *services.TokenService
}

func (s *apiServer) beforeStart() {
	s.UserStore = services.NewUserStore(s.PgConnection)
	s.TokenService = services.NewTokenStore(s.RedisClient)

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

func (s *apiServer) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	data := map[string]string{"error": err.Error()}
	s.respond(w, r, code, data)
}

func (s *apiServer) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

func NewServer(logger *logrus.Logger, config *Config, pqConnection *sql.DB, redisClient *redis.Client) *apiServer {
	return &apiServer{
		logger:       logger,
		Config:       config,
		router:       mux.NewRouter(),
		PgConnection: pqConnection,
		RedisClient:  redisClient,
	}
}
