package controller

import (
	"github.com/gin-gonic/gin"
	"keycloak/keycloak"
)

type Server struct {
	router   *gin.Engine
	keycloak *keycloak.Keycloak
}

func NewServer(kc *keycloak.Keycloak) *Server {
	server := &Server{
		keycloak: kc,
	}
	server.setupRouter()
	return server
}

func (server *Server) setupRouter() {
	router := gin.Default()

	//routes
	router.POST("/login", server.login)

	//routes that require authentication
	authRoute := router.Group("/").Use(authMiddleware(server.keycloak))
	authRoute.GET("/docs", server.getDocs)

	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
