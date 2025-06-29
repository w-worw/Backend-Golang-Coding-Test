package middleware

import (
	"7-solutions/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization token"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token format"})
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization token"})
			c.Abort()
			return
		}

		email, ok := claims["email"]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "email not found in token"})
			c.Abort()
			return
		}

		name, ok := claims["name"]
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "name not found in token"})
			c.Abort()
			return
		}

		c.Set("email", email)
		c.Set("name", name)
		c.Next()
	}
}
