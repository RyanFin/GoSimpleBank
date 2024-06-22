package gapi

import (
	db "RyanFin/GoSimpleBank/db/sqlc"
	"RyanFin/GoSimpleBank/pb"
	"RyanFin/GoSimpleBank/token"
	"RyanFin/GoSimpleBank/util"
	"fmt"
)

// Server serves gRPC requests for our banking service
type Server struct {
	pb.UnimplementedSimpleBankServer
	config     util.Config
	store      db.Store // an interface instead of a struct
	tokenMaker token.Maker
}

// NewServer creates a new gRPC server and sets up routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("unable to instantiate token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// No routes in gRPC no set server router method required
	// server.setupRouter()

	return server, nil
}
