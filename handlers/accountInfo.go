package handlers

import "github.com/gin-gonic/gin"

func AccountInfo(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(200, user)
}
