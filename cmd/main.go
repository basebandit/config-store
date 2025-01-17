package main

import (
	"log"
	"net/http"
	"os"

	configstore "github.com/basebandit/config-store"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultHTTPPort = "3000"

func main() {
	// Load environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultHTTPPort
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is required")
	}

	// Initialize the database
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	// Auto-migrate the KV model
	if err := db.AutoMigrate(&configstore.KV{}); err != nil {
		log.Fatalf("Failed to auto-migrate database: %v", err)
	}

	// Initialize KVService
	kvService := configstore.NewKVService(db)

	// Initialize router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorizzation", "Content-Type"},
	}))

	// Routes
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("config-store!"))
	})

	r.Mount("/api", configstore.ApiRouter(kvService))

	// Start http server
	log.Printf("Server running on port %s", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
