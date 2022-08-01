package services

import (
	"github.com/go-redis/redis"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/maxsnegir/rest-api-go/internal/app/models"
	"time"
)

var (
	jwtSecret        = []byte("superSecretKey")
	AccessTokenType  = "access"
	RefreshTokenType = "refresh"
)

type TokenService struct {
	store *redis.Client
}

type JWTClaim struct {
	TokenType string `json:"token_type"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	jwt.RegisteredClaims
}

func (t *TokenService) SetToken(identity, token string) error {
	err := t.store.Set(identity, token, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func (t *TokenService) getToken(identity string) (string, error) {
	result, err := t.store.Get(identity).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}

func (t *TokenService) deleteToken(identity string) {
	t.store.Del(identity)
}
func (t *TokenService) CreateAccessToken(user models.User) (string, error) {

	claims := &JWTClaim{
		AccessTokenType,
		user.Username,
		user.Email,
		jwt.RegisteredClaims{
			Subject:   user.Id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 5)),
		},
	}
	return createToken(claims)
}

func (t *TokenService) CreateRefreshToken(user models.User) (string, error) {
	claims := &JWTClaim{
		RefreshTokenType,
		user.Username,
		user.Email,
		jwt.RegisteredClaims{
			Subject:   user.Id,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			ID:        uuid.New().String(),
		},
	}
	return createToken(claims)
}

func (t *TokenService) CreateTokens(user models.User) (*models.Token, error) {
	accessToken, err := t.CreateAccessToken(user)
	if err != nil {
		return nil, err
	}
	refreshToken, err := t.CreateRefreshToken(user)
	if err != nil {
		return nil, err
	}
	tokens := &models.Token{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	return tokens, nil
}

func createToken(claims *JWTClaim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func NewTokenStore(store *redis.Client) *TokenService {
	return &TokenService{
		store: store,
	}
}
