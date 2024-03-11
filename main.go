package main

import (
	"backend/config"
	"backend/database"
	"backend/handlers"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.InfoLevel)

	// Load configuration
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	// Setup router
	router := handlers.SetupRouter(db)

	// Start server
	port := "9080" // You can change the port as per your requirement
	fmt.Printf("Server is running on port %s...\n", port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
