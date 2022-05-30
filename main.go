package main

import (
	"database/sql"
	"log"

	docs "github.com/STAMBOULI-ABDELKARIM/car_repair_shop/docs"
	_ "github.com/lib/pq"

	db "github.com/STAMBOULI-ABDELKARIM/car_repair_shop/db/sqlc"

	"github.com/STAMBOULI-ABDELKARIM/car_repair_shop/api"
	"github.com/STAMBOULI-ABDELKARIM/car_repair_shop/util"
)

func main() {
	docs.SwaggerInfo.Title = "ABDELKARIM STAMBOULI - CAR REPAIR SHOP API"
	docs.SwaggerInfo.Description = "MANAGE CUSTOMERS FOR A CAR REPAR SHOPAPI."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
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
