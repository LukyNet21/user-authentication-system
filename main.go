package main

import (
	"auth-system/handlers"
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

	r := gin.Default()

	v1 := r.Group("/api/v1/")
	v1.POST("/login", handlers.Login)
	v1.POST("/register", handlers.Register)
	r.Run(":8080")
}
