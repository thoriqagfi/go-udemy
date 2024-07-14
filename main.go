package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"go-udemy.sqlc.dev/app/api"
	db "go-udemy.sqlc.dev/app/db/sqlc"
	"go-udemy.sqlc.dev/app/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		panic(err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		panic(err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		panic(err)
	}
}