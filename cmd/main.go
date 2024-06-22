package main

import (
	"log"

	"github.com/RobTov/hmblog-golang-backend/cmd/api"
	"github.com/RobTov/hmblog-golang-backend/config"
	"github.com/RobTov/hmblog-golang-backend/db"
)

func main() {
	db, err := db.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":"+config.Envs.Port, db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
