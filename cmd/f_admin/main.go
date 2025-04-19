package main

import (
	"f_admin_go/internal/api/assets"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/config"
	"f_admin_go/internal/db"
	"log"
	"net/http"
	"time"
)

func main() {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	cfg := config.LoadConfig()

	log.Printf("Server running on port %s", cfg.Port)
	log.Printf("Environment: %s", cfg.Environment)
	log.Println("Connecting to database...")

	if err := db.InitDB(cfg.DBURL); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connection successful")
	}
	defer func() {
		if err := db.DB.Close(); err != nil {
			log.Fatalf("Failed to close database connection: %v", err)
		}
		log.Println("Database connection closed")
	}()

	// TODO: Add healthcheck cycle

	log.Fatal(server.ListenAndServe())
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/assets", shared.HandleEntity(&assets.Handler{}))
	// mux.HandleFunc("/api/product", api.HandleProductRequest)
	// mux.HandleFunc("/api/transaction", api.HandleTransactionRequest)
	mux.HandleFunc("/", handleNotFound)
	return mux
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to fantasy Tech"))
}
