package middleware

import (
	"api/internal/models"
	"api/pkg/db"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RequireAuth(c *gin.Context) {
	log.Println("Validating token")

	hmacSampleSecret := []byte(os.Getenv("JWT_SECRET"))

	tokenString, err := c.Cookie("Auth")
	if err != nil {
		log.Println("Error retrieving token:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// parse token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			log.Printf("Unexpected signing method: %v", token.Header["alg"])
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Sprawdzenie, czy token jest ważny
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp, ok := claims["expire"].(float64)
		if !ok || time.Now().Unix() > int64(exp) {
			log.Println("Token expired or invalid")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			log.Println("Invalid token subject")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		var user models.User
		if err := db.DB.First(&user, int(sub)).Error; err != nil {
			log.Println("Error retrieving user from database:", err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.ID == 0 {
			log.Println("User not found")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// set user in context
		c.Set("user", user)
		c.Next()

	} else {
		log.Println("Invalid token claims")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
