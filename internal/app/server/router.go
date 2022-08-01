package server

import "net/http"

func (s *apiServer) configureRouter() {
	s.router.Use(loggingMiddleware)
	s.router.HandleFunc("/sign-up/", s.handleSignUp()).Methods(http.MethodPost)
	s.router.HandleFunc("/sign-in/", s.handleSignIn()).Methods(http.MethodPost)
	s.router.Handle("/users/me/", IsAuthenticated(s.handleGetMe, false)).Methods(http.MethodGet)

}
