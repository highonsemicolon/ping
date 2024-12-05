package middleware

import (
	"fmt"
	"net/http"

	"user-service/pkg/utils"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gin-gonic/gin"
)

func JwtMiddleware(config *utils.Auth0) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {

				checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer("https://"+config.Domain+"/", false)
				if !checkIss {
					return token, fmt.Errorf("Invalid issuer")
				}

				audiences := token.Claims.(jwt.MapClaims)["aud"].([]interface{})
				checkAud := false
				for _, audience := range audiences {
					if audience == config.Audience {
						checkAud = true
						break
					}
				}

				if !checkAud {
					return token, fmt.Errorf("Invalid audience")
				}

				// Get the public key for RS256
				cert, err := utils.GetPemCert(config.Domain)
				if err != nil {
					return nil, err
				}

				result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
				return result, nil
			},
			SigningMethod: jwt.SigningMethodRS256,
		})

		err := jwtMiddleware.CheckJWT(c.Writer, c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}
