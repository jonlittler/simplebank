package main

import (
	"database/sql"
	"log"

	"github.com/jonlittler/ts/simplebank/api"
	db "github.com/jonlittler/ts/simplebank/db/sqlc"
	"github.com/jonlittler/ts/simplebank/util"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatalf("cannot create server: %v", err)
	}

	if err := server.Start(config.ServerAddr); err != nil {
		log.Fatalf("cannot start http server: %v", err)
	}
}
