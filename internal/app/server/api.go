package server

import (
	"encoding/json"
	"github.com/maxsnegir/rest-api-go/internal/app/models"
	"net/http"
)

func (s *apiServer) handleSignUp() http.HandlerFunc {
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
		tokens, err := s.TokenService.CreateTokens(*user)
		if err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		s.respond(w, r, http.StatusCreated, tokens)
	}
}

func (s *apiServer) handleSignIn() http.HandlerFunc {
	type request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		user, err := s.UserStore.GetByUsername(req.Username)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		tokens, err := s.TokenService.CreateTokens(*user)
		if err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		if err := s.TokenService.SetToken(user.Id, tokens.RefreshToken); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, tokens)
	}
}

func (s *apiServer) handlerSignOut() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *apiServer) handleGetMe(w http.ResponseWriter, r *http.Request) {

	requestUser, ok := r.Context().Value("requestUser").(RequestUser)

	if !ok {
		return
	}
	user, err := s.UserStore.GetByUsername(requestUser.Username)
	if err != nil {
		return
	}
	s.respond(w, r, http.StatusOK, user)
}
