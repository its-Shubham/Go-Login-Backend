package handlers

import (
	"backend/repository"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	// Initialize Gin router
	router := gin.Default()

	// Initialize repository
	userRepo := repository.NewUserRepository(db)

	// Define routes
	v1 := router.Group("/api/v1")
	{
		// User routes
		v1.POST("/register", RegisterHandler(userRepo))
		v1.POST("/login", LoginHandler(userRepo))
		v1.GET("/users", getUserList(userRepo))
		// Add more routes here as needed
	}

	return router
}
