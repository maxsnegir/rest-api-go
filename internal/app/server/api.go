package server

import (
	"io"
	"net/http"
)

func (s *apiServer) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := io.WriteString(w, "HELLO"); err != nil {
			s.logger.Error(err)
		}
	}
}
