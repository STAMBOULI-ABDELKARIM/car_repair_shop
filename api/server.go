package api

import (
	db "github.com/STAMBOULI-ABDELKARIM/car_repair_shop/db/sqlc"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Server struct {
	store  *db.Queries
	router *gin.Engine
}

func NewServer(store *db.Queries) *Server {
	server := &Server{store: store}
	router := gin.Default()
	server.router = router

	router.GET("/customers/:id", server.getCustomer)
	router.POST("/customers", server.createCustomer)
	router.PUT("/customers/:id", server.updateCustomer)
	router.DELETE("/customers/:id", server.deleteCustomer)
	router.GET("/customers", server.listCustomers)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return server
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
