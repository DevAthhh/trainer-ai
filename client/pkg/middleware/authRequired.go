package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DevAthhh/trainer-ai/client/pkg/initializers"
	"github.com/DevAthhh/trainer-ai/client/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorizated",
		})
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user models.UserTrainers
		initializers.DB.First(&user, claims["sub"])
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Unauthorizated",
			})
		}

		c.Set("user", user)

		c.Next()

	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorizated",
		})
	}
}
