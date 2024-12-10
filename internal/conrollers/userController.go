package conrollers

import (
	"bar_site/configs"
	"bar_site/internal/models"
	"bar_site/internal/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func SignUp(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) == nil {
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

	c.SetCookie("access_token", tokenString, 3600, "/", "", true, true)
}
