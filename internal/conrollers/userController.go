package conrollers

import (
	"bar_site/configs"
	"bar_site/internal/models"
	"bar_site/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to generate password hash"})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := configs.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create user"})
		return
	}

	tokenString, err := utils.CreateJWTToken(user.ID, c.ClientIP())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create JWT"})
	}
	refreshToken, err := utils.GenerateRefreshToken()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create refresh token"})
	}
	hashRefresh, err := utils.HashRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create refresh token"})
	}

	user.RefreshToken = hashRefresh
	configs.DB.Save(&user)

	c.SetCookie("access_token", tokenString, 3600, "/", "", true, false)
	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func SignIn(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read body"})
		return
	}

	var user models.User
	if err := configs.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Неверный email или пароль"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	if err := utils.CheckPasswordHash(body.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}

	token, err := utils.CreateJWTToken(user.ID, c.ClientIP())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create JWT"})
	}

	c.SetCookie("access_token", token, 3600, "/", "", true, false)
	c.JSON(http.StatusOK, gin.H{"message": "User auth successfully", "token": token})
}
