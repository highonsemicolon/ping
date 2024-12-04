package api

import (
	"net/http"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProfileHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims := c.Request.Context().Value("user").(*jwt.Token).Claims.(jwt.MapClaims)

		username := claims["username"].(string)
		email := claims["email"].(string)

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"email":    email,
		})
	}
}
