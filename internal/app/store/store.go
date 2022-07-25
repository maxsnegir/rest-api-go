package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Store interface {
	Connect() (*sql.DB, error)
	Close() error
}

type store struct {
	Connection *sql.DB
	driver     string
	dbUrl      string
}

func (s *store) Connect() error {
	connection, err := sql.Open(s.driver, s.dbUrl)
	if err != nil {
		return err
	}

	if err := connection.Ping(); err != nil {
		return err
	}

	s.Connection = connection
	return nil
}

func (s *store) Close() (err error) {
	err = s.Connection.Close()
	if err != nil {
		return
	}
	fmt.Printf("Connection to %s is closed\n", s.driver)
	return
}

func NewStore(driver, dbUrl string) *store {
	return &store{
		driver: driver,
		dbUrl:  dbUrl,
	}
}
