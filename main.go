package main

import (
	"RyanFin/GoSimpleBank/api"
	db "RyanFin/GoSimpleBank/db/sqlc"
	"RyanFin/GoSimpleBank/gapi"
	"RyanFin/GoSimpleBank/pb"
	"RyanFin/GoSimpleBank/util"
	"database/sql"
	"log"
	"net"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	// runHTTPServer(config, store)

	runGrpcServer(config, store)
}

// gRPC Server
func runGrpcServer(config util.Config, store db.Store) {
	server, err := gapi.NewServer(config, store)
	if err != nil {
		log.Fatal("instantiate new server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterSimpleBankServer(grpcServer, server)

	// powerful statement that allows grpc client to easily explore available RPCs on the server
	// and call them
	reflection.Register(grpcServer)

	// start the server
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("cannot create listener")
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server listener")
	}
}

// Gin HTTP Server
func runHTTPServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("instantiate new server:", err)
	}

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal("cannot start HTTP server: ", err)
	}
}
