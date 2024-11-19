package middleware

import (
	"auth-system/models"
	"auth-system/utils"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil || !token.Valid {
		c.JSON(401, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if exp, ok := claims["exp"].(float64); ok && float64(time.Now().Unix()) > exp {
			c.JSON(401, gin.H{"error": "Token expired"})
			c.Abort()
			return
		}
		var user models.User
		err := utils.DB.Where("id = ?", claims["user_id"]).First(&user).Error
		if err != nil {
			c.JSON(401, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		user.Password = ""
		c.Set("user", user)
		c.Next()
		return
	}

	fmt.Println(err)
	c.JSON(401, gin.H{"error": "Unauthorized"})
	c.Abort()

}
