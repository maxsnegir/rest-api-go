package store

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type store struct {
	Connection *sql.DB
	driver     string
	dbUrl      string
}

func (s *store) Connect() error {
	db, err := sql.Open(s.driver, s.dbUrl)
	if err != nil {
		return err
	}

	if err := db.Ping(); err != nil {
		return err
	}

	s.Connection = db
	fmt.Printf("Сonnected to %s\n", s.driver)
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
