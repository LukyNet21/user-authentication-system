package main

import (
	"auth-system/handlers"
	"auth-system/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.InitDB()

	r := gin.Default()

	v1 := r.Group("/api/v1/")
	v1.POST("/login", handlers.Login)
	v1.POST("/register", handlers.Register)
	r.Run(":8080")
}
