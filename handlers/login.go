package handlers

import (
	"auth-system/models"
	"auth-system/utils"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var loginUser models.User
	var dbUser models.User

	if err := c.BindJSON(&loginUser); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON format"})
		return
	}

	// Check if user exists
	err := utils.DB.Where("username = ?", loginUser.Username).First(&dbUser).Error
	if err != nil {
		c.JSON(400, gin.H{"error": "User not found"})
		return
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(loginUser.Password))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid password"})
		return
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  dbUser.ID,
		"username": dbUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
		"iat":      time.Now().Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		c.JSON(500, gin.H{"error": "JWT_SECRET environment variable not set"})
		return
	}

	token, err := claims.SignedString([]byte(jwtSecret))
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token"})
		return
	}

	serverAddress := os.Getenv("SERVER_ADDRESS")
	if serverAddress == "" {
		c.JSON(500, gin.H{"error": "SERVER_ADDRESS environment variable not set"})
		return
	}
	c.SetCookie("token", token, 3600*24, "/", serverAddress, false, true)

	c.JSON(200, gin.H{"message": "succsessfully logged in"})
}
