package controller

import (
	"fmt"
	"net/http"
	"time"
	"tukerin/config"
	"tukerin/models"
	"tukerin/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {

	var user models.User
	var cart models.Cart

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.Email == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and Password are required"})
		return
	}
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name is required"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hash)

	if config.DB.Create(&user).Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	cart.UserId = int(user.ID)
	
	if err := config.DB.Create(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

func Login(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	if err := config.DB.Where("email = ?", user.Email).First(&existingUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Generate OTP
	otp, err := utils.GenerateOTP(existingUser.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate OTP"})
		return
	}

	var expires_at = time.Now().Add(5 * time.Minute) 

	otpModel := models.Otp{
		UserId: int(existingUser.ID),
		Email:  existingUser.Email,
		Code:   otp,
		Type:   "login",
		ExpiresAt: expires_at, // OTP valid for 5 minutes
	}

	if err := config.DB.Create(&otpModel).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create OTP"})
		return
	}

	utils.SendEmail(existingUser.Email, "Your OTP Code", "Your OTP code is: "+otp)

	c.JSON(http.StatusOK, gin.H{"message": "OTP sent to your email"})

}

func VerifyOTP(c *gin.Context) {
	var request struct {
		Email string `json:"email" form:"email"`
		Code  string `json:"otp" form:"otp"`
	}

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var otp models.Otp
	var user models.User

	fmt.Println("Verifying OTP for email:", request.Email)

	if err := config.DB.Where("email = ?", request.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	if err := config.DB.Where("user_id = ? AND code = ? AND expires_at > NOW()", user.ID, request.Code).Order("created_at DESC").First(&otp).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid OTP"})
		return
	}

	token, err := utils.GenerateJWT(user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	
	if err := config.DB.Delete(&otp).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete OTP"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "OTP verified successfully", "token": token})
}
