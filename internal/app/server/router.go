package server

import "net/http"

func (s *apiServer) configureRouter() {
	s.router.Use(loggingMiddleware)
	s.router.HandleFunc("/user/create", s.handleCreateUser()).Methods(http.MethodPost)
}
