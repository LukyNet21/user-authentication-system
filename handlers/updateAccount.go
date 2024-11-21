package handlers

import (
	"auth-system/models"
	"auth-system/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

func UpdateAccount(c *gin.Context) {
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

	var userUpdated models.User
	if err := c.BindJSON(&userUpdated); err != nil {
		c.JSON(400, gin.H{"error": "Invalid form"})
		return
	}

	if userUpdated.Username != "" {
		c.JSON(400, gin.H{"error": "Username cannot be updated"})
		return
	}

	validate := validator.New()
	if userUpdated.Email != "" {
		if err := validate.Var(userUpdated.Email, "email"); err != nil {
			c.JSON(400, gin.H{"error": "Invalid email"})
			return
		}
		user.Email = userUpdated.Email
	}

	if userUpdated.Password != "" {
		if err := validate.Var(userUpdated.Password, "min=8"); err != nil {
			c.JSON(400, gin.H{"error": "Password must be at least 8 characters long"})
			return
		}

		// Hash password
		bytes, err := bcrypt.GenerateFromPassword([]byte(userUpdated.Password), 14)
		if err != nil {
			c.JSON(500, gin.H{"message": "internal server error"})
			return
		}
		user.Password = string(bytes)
	}

	if userUpdated.FirstName != "" {
		user.FirstName = userUpdated.FirstName
	}

	if userUpdated.LastName != "" {
		user.LastName = userUpdated.LastName
	}

	err = utils.DB.Save(&user).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(200, gin.H{"message": "Account updated successfully"})
}
