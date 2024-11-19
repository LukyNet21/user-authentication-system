package handlers

import (
	"auth-system/models"
	"auth-system/utils"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
)

func ResetPassword(c *gin.Context) {
	var json struct {
		Email string `json:"email"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON format"})
		return
	}

	var user models.User
	if err := utils.DB.Where("email = ?", json.Email).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	randomString := utils.RandString(256)
	resetPassword := models.ResetPassword{
		UserId:  user.ID,
		Token:   randomString,
		ValidTo: time.Now().Add(time.Hour),
	}

	if err := utils.DB.Create(&resetPassword).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create reset password token"})
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL_SERVER_USERNAME"))
	m.SetHeader("To", json.Email)
	m.SetHeader("Subject", "Reset password")
	m.SetBody("text/html", "Click <a href='"+os.Getenv("RESET_PASSWORD_URL")+randomString+"'>here</a> to reset your password. This link is valid for 1 hour.")

	if err := utils.Email.DialAndSend(m); err != nil {
		c.JSON(500, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(200, gin.H{"mail": "reset email sent"})
}

func ResetPasswordToken(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.JSON(400, gin.H{"error": "Token is required"})
		return
	}

	var resetPassword models.ResetPassword
	if err := utils.DB.Where("token = ?", token).First(&resetPassword).Error; err != nil {
		c.JSON(404, gin.H{"error": "Token not found"})
		return
	}

	if time.Now().After(resetPassword.ValidTo) {
		c.JSON(400, gin.H{"error": "Token expired"})
		return
	}

	var user models.User
	if err := utils.DB.Where("id = ?", resetPassword.UserId).First(&user).Error; err != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	var json struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(400, gin.H{"error": "invalid JSON format"})
		return
	}
	validate := validator.New()
	if err := validate.Var(user.Password, "required,min=8"); err != nil {
		c.JSON(400, gin.H{"message": "password must be at least 8 characters"})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(json.Password), 14)
	if err != nil {
		c.JSON(500, gin.H{"message": "internal server error"})
		return
	}
	user.Password = string(bytes)

	if err := utils.DB.Save(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(200, gin.H{"message": "Password updated successfully"})

	utils.DB.Where("token = ?", token).Delete(&models.ResetPassword{})
}
