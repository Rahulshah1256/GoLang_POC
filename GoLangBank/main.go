package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/techschool/simplebank/api"
	"github.com/techschool/simplebank/util"
	"golangbank/docs"

	db "github.com/techschool/simplebank/db/sqlc"
)

// @title 	Simple Bank Service API
// @version	1.0
// @description A Simple Bank service API in Go using Gin framework

// @host 	localhost:8888
// @BasePath /api
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, *store)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
