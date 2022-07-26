package server

import (
	"encoding/json"
	"github.com/maxsnegir/rest-api-go/internal/app/models"
	"net/http"
)

func (s *apiServer) handleCreateUser() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		user := &models.User{
			Username: req.Username,
			Email:    req.Email,
			Password: req.Password,
		}
		if err := s.UserStore.Create(user); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, user)
	}
}
