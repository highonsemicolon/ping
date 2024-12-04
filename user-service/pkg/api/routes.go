package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/highonsemicolon/ping/user-service/pkg/middleware"
	"github.com/highonsemicolon/ping/user-service/pkg/utils"
)

func SetupRoutes(router *gin.Engine, db *sql.DB, config utils.Auth0) {
	router.Use(cors.New(cors.Config{AllowAllOrigins: true}))

	router.GET("/", PublicEndpointHandler)

	protected := router.Group("/")
	protected.Use(middleware.JwtMiddleware(&config))
	protected.GET("/protected-endpoint", ProtectedEndpointHandler)
	protected.GET("/profile", ProfileHandler())
}

func PublicEndpointHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Public endpoint!"})
}

func ProtectedEndpointHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You are authorized!"})
}
