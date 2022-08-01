package services

import (
	"database/sql"
	"github.com/maxsnegir/rest-api-go/internal/app/models"
)

type UserStore interface {
	Create(u *models.User) error
	Update()
	Delete()
	GetByUsername(username string) (*models.User, error)
}

type userStore struct {
	connection *sql.DB
}

func NewUserStore(connection *sql.DB) *userStore {
	return &userStore{connection: connection}
}

func (us *userStore) Create(user *models.User) error {
	if err := user.Validate(); err != nil {
		return err
	}
	row := us.connection.QueryRow(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Username,
		user.Email,
		user.Password,
	)
	if err := row.Scan(&user.Id); err != nil {
		return err
	}
	return nil
}

func (us *userStore) GetByUsername(username string) (*models.User, error) {
	user := &models.User{}
	err := us.connection.QueryRow(
		"SELECT id, username, email, password from users WHERE username = $1",
		username,
	).Scan(
		&user.Id,
		&user.Username,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (us *userStore) Update() {

}

func (us *userStore) Delete() {

}
