package main

import (
	"database/sql"
	"github.com/CeciliaChen/simplebank/api" // Replace with the actual import path for api
	db "github.com/CeciliaChen/simplebank/db/sqlc"
	_ "github.com/lib/pq" // This is the driver for PostgreSQL
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
