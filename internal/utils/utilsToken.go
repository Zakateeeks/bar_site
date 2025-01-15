package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"os"
	"time"
)

// GenerateRefreshToken Функция для генерации refresh токена
func GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// CreateJWTToken Функция для создания access токена
func CreateJWTToken(userID uint, clientIP string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
		"ip":  clientIP,
	})

	return token.SignedString([]byte(os.Getenv("TOKEN")))
}

// HashRefreshToken Функция для хэширования токена
func HashRefreshToken(refreshToken string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(refreshToken), 12)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func ValidateJWTToken(tokenString string) (*jwt.Token, error) {
	Secret := []byte(os.Getenv("TOKEN"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return Secret, nil
	})

	if err != nil {
		return nil, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return token, nil
	}

	return nil, errors.New("invalid token")
}

func CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Println("Password does not match")
		return err
	}
	return nil
}
