package store

import (
	"database/sql"
	"github.com/maxsnegir/rest-api-go/internal/app/models"
)

type UserStore interface {
	Create(u *models.User) error
	Update()
	Delete()
}

type userStore struct {
	connection *sql.DB
}

func NewUserStore(connection *sql.DB) *userStore {
	return &userStore{connection: connection}
}

func (us *userStore) Create(user *models.User) error {
	row := us.connection.QueryRow(
		"INSERT INTO users (username, email, password) VALUES ($1, $2, $3) RETURNING id",
		user.Username,
		user.Email,
		user.Password,
	)
	if err := row.Scan(&user); err != nil {
		return err
	}
	return nil
}

func (us *userStore) Update() {

}

func (us *userStore) Delete() {

}
