package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/jonlittler/ts/simplebank/db/sqlc"
	"github.com/jonlittler/ts/simplebank/token"
	"github.com/jonlittler/ts/simplebank/util"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	config     util.Config
	store      db.Store
	router     *gin.Engine
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := Server{config: config, store: store, tokenMaker: tokenMaker}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return &server, nil
}

// Start runs the HTTP server on a specific address.
func (s *Server) Start(addr string) error {
	return s.router.Run(addr)
}

// Error handler for gin returns gin.H (map[string]any).
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (s *Server) setupRouter() {

	/* create router */
	router := gin.Default()

	router.POST("/accounts", s.createAccount)

	/* add routes to router */
	/* open routes */
	// router.POST("/users", s.createUser)
	// router.POST("/users/login", s.loginUser)

	// /* protected routes */
	// authRoutes := router.Group("/")
	// authRoutes.Use(authMiddleware(s.tokenMaker))

	// authRoutes.POST("/accounts", s.createAccount)
	// authRoutes.GET("/accounts/:id", s.getAccountByID)
	// authRoutes.GET("/accounts", s.listAccounts)
	// authRoutes.PUT("/accounts/:id", s.updateAccountByID)
	// authRoutes.DELETE("/accounts/:id", s.deleteAccountByID)

	// authRoutes.POST("/transfers", s.createTransfer)

	/* save router to store */
	s.router = router
}
