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
	account := v1.Group("/account/")
	account.POST("/login", handlers.Login)
	account.POST("/register", handlers.Register)
	account.POST("/reset-password", handlers.ResetPassword)
	account.PUT("/reset-password/:token", handlers.ResetPasswordToken)
	account.PATCH("/update/:id", middleware.AuthMiddleware, handlers.UpdateAccount)
	account.DELETE("/delete/:id", middleware.AuthMiddleware, handlers.DeleteAccount)
	account.GET("/info", middleware.AuthMiddleware, handlers.AccountInfo)
	r.Run(":8080")
}
