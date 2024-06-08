package api

import (
	db "RyanFin/GoSimpleBank/db/sqlc"
	"RyanFin/GoSimpleBank/token"
	"RyanFin/GoSimpleBank/util"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      db.Store // an interface instead of a struct
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing
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

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()
	// Set trusted proxies to localhost
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		panic(err)
	}
	// anyone should have access to these endpoints
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	// apply the authMiddleware to all the endpoints below
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	// add routes to the router with handlers
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccounts)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// run HTTP server on specific address to listen for requests
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
