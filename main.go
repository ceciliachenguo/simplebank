package main

import (
	"database/sql"
	"log"

	"github.com/CeciliaChen/simplebank/api" // Replace with the actual import path for api
	db "github.com/CeciliaChen/simplebank/db/sqlc"
	"github.com/CeciliaChen/simplebank/util"
	_ "github.com/lib/pq" // This is the driver for PostgreSQL
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
