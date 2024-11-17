package handlers

import (
	"auth-system/models"
	"auth-system/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func Register(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "invalid JSON format"})
		return
	}

	if user.Username == "" || user.Password == "" || user.Email == "" || user.FirstName == "" || user.LastName == "" {
		c.JSON(400, gin.H{"message": "missing required fields"})
		return
	}

	validate := validator.New()
	if err := validate.Var(user.Username, "required,min=3,max=30"); err != nil {
		c.JSON(400, gin.H{"message": "username must be between 3 and 30 characters"})
		return
	}

	if err := validate.Var(user.Email, "required,email"); err != nil {
		c.JSON(400, gin.H{"message": "invalid email format"})
		return
	}

	if err := validate.Var(user.Password, "required,min=8"); err != nil {
		c.JSON(400, gin.H{"message": "password must be at least 8 characters"})
		return
	}

	err := utils.DB.Create(&user).Error
	if err != nil {
		if err == gorm.ErrDuplicatedKey {
			c.JSON(400, gin.H{"message": "username or email already exists"})
			return
		}
		c.JSON(500, gin.H{"message": "internal server error"})
		return
	}
	c.JSON(200, gin.H{"message": "user created successfully"})
}
