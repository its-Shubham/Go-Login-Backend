package handlers

import (
	"backend/models"
	"backend/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandler(userRepo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := userRepo.Register(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
			return
		}
		c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
	}
}

// LoginHandler wraps the userRepo.Login method and serves as a Gin handler for user login.
func LoginHandler(userRepo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginReq models.LoginRequest
		if err := c.ShouldBindJSON(&loginReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userRepo.Login(loginReq.Email, loginReq.Password)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusOK, user)
	}
}
