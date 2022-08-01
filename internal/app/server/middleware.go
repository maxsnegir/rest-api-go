package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/maxsnegir/rest-api-go/internal/app/services"
	"log"
	"net/http"
	"strings"
)

type RequestUser struct {
	IsAuthenticated bool
	Id              string
	Username        string
	Email           string
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s :: %s ", r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}

var jwtSecret = []byte("superSecretKey")

func IsAuthenticated(endpoint func(w http.ResponseWriter, r *http.Request), refresh bool) http.Handler {

	requestUser := RequestUser{}
	var tokenType string
	switch refresh {
	case true:
		tokenType = services.RefreshTokenType
	default:
		tokenType = services.AccessTokenType
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := getBearerToken(r.Header)
		if err != nil {
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &services.JWTClaim{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			claims, ok := token.Claims.(*services.JWTClaim)
			if !ok {
				return nil, fmt.Errorf("wrong token")
			}

			if err := claims.Valid(); err != nil {
				return nil, err
			}
			if claims.TokenType != tokenType {
				return nil, fmt.Errorf("wrong token type")
			}
			setRequestUser(&requestUser, claims)
			return jwtSecret, nil
		})
		ctx := r.Context()
		req := r.WithContext(context.WithValue(ctx, "requestUser", requestUser))
		if err == nil && token.Valid {
			endpoint(w, req)
		}

	})
}

func getBearerToken(header http.Header) (string, error) {
	authToken, ok := header["Authorization"]
	if !ok {
		return "", errors.New("there is no Authorization header")
	}
	tokenParts := strings.Split(authToken[0], " ") // Bearer token
	if tokenParts[0] != "Bearer" {
		return "", errors.New("it's not Bearer token")
	}
	return tokenParts[1], nil
}

func setRequestUser(requestUser *RequestUser, claims *services.JWTClaim) {
	requestUser.IsAuthenticated = true
	requestUser.Id = claims.Subject
	requestUser.Username = claims.Username
	requestUser.Email = claims.Email
}
