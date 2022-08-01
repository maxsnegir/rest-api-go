package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type pqStore struct {
	Connection *sql.DB
	DbDsn      string
}

func (pq *pqStore) Connect() error {
	connection, err := sql.Open("postgres", pq.DbDsn)
	if err != nil {
		return err
	}

	if err := connection.Ping(); err != nil {
		return err
	}

	pq.Connection = connection
	return nil
}

func (pq *pqStore) Close() (err error) {
	err = pq.Connection.Close()
	if err != nil {
		return
	}
	fmt.Println("connection to postgres is closed")
	return
}

func NewPqStore(DbDsn string) *pqStore {
	return &pqStore{
		DbDsn: DbDsn,
	}
}
