package main

import (
	"RyanFin/GoSimpleBank/api"
	db "RyanFin/GoSimpleBank/db/sqlc"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@localhost:5432/simple_bank?sslmode=disable"
	addr     = "localhost:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to the DB:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(addr)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
