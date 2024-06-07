package main

import (
	"RyanFin/GoSimpleBank/api"
	db "RyanFin/GoSimpleBank/db/sqlc"
	"RyanFin/GoSimpleBank/util"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// load config with viper. app.yml is located in the same directory as main.go in this instance
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to the DB:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("instantiate new store:", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
