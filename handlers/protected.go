package handlers

import "github.com/gin-gonic/gin"

func Protected(c *gin.Context) {
	c.JSON(200, gin.H{"data": "protected", "user": c.MustGet("user")})
}
