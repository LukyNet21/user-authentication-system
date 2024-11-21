package handlers

import (
	"auth-system/models"
	"auth-system/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	user := c.MustGet("user").(models.User)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
		return
	}
	if int(user.ID) != id {
		c.JSON(403, gin.H{"error": "Forbidden"})
		return
	}

	err = utils.DB.Delete(&user).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(200, gin.H{"message": "Account deleted successfully"})

}
