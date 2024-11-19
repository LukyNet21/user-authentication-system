package main

import (
	"auth-system/handlers"
	"auth-system/middleware"
	"auth-system/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	utils.InitDB()
	utils.InitEmail()
	// Send a test email
	// utils.SendTestEmail("mail@example.com")

	r := gin.Default()

	v1 := r.Group("/api/v1/")
	v1.POST("/login", handlers.Login)
	v1.POST("/register", handlers.Register)
	v1.POST("/reset-password", handlers.ResetPassword)
	v1.PUT("/reset-password/:token", handlers.ResetPasswordToken)
	v1.GET("/protected", middleware.AuthMiddleware, handlers.Protected)
	r.Run(":8080")
}
