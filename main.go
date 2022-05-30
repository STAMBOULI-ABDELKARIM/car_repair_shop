package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	db "github.com/STAMBOULI-ABDELKARIM/car_repair_shop/db/sqlc"

	"github.com/STAMBOULI-ABDELKARIM/car_repair_shop/api"
	"github.com/STAMBOULI-ABDELKARIM/car_repair_shop/util"
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

	store := db.New(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start the server", err)
	}

}
