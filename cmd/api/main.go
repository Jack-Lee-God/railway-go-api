package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"goapi.railway.app/internal/database"
	"goapi.railway.app/internal/models"
	"gorm.io/gorm"
)

const version = "0.0.1"

type config struct {
	port int
}

type application struct {
	config config
	db     *gorm.DB
}

func main() {

	var cfg config

	// Try to read environment variable for port (given by railway). Otherwise use default
	port := os.Getenv("PORT")
	intPort, err := strconv.Atoi(port)
	if err != nil {
		intPort = 4000
	}

	// Set the port to run the API on
	cfg.port = intPort

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	} else {
		log.Println("Connected to database successfully")
		// Auto migrate models
		err = db.AutoMigrate(&models.Bioskop{})
		if err != nil {
			log.Printf("Failed to migrate database: %v", err)
		}
	}

	// create the application
	app := &application{
		config: cfg,
		db:     db,
	}

	// create the server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  45 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("server started on %s", srv.Addr)

	// Start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
