package api

import (
	"database/sql"
	"net/http"

	"user-service/pkg/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, db *sql.DB, domain, audience string) {
	router.Use(cors.New(cors.Config{AllowAllOrigins: true}))

	router.GET("/", PublicEndpointHandler)

	protected := router.Group("/")
	protected.Use(middleware.JwtMiddleware(domain, audience))
	protected.GET("/protected-endpoint", ProtectedEndpointHandler)
	protected.GET("/profile", ProfileHandler())
}

func PublicEndpointHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Public endpoint!"})
}

func ProtectedEndpointHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
}
