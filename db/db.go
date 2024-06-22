package db

import (
	"database/sql"
	"log"

	"github.com/RobTov/hmblog-golang-backend/config"
	_ "github.com/lib/pq"
)

func NewPostgresStorage() (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Envs.DatabaseURI)
	if err != nil {
		log.Fatal(err)
	}

	err = initStorage(db)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("DB successfully connected\n")

	return db, nil
}

func initStorage(db *sql.DB) error {
	return db.Ping()
}
