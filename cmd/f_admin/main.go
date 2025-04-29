package main

import (
	"f_admin_go/internal/api/assets"
	"f_admin_go/internal/api/auth"
	"f_admin_go/internal/api/membership"
	"f_admin_go/internal/api/products"
	"f_admin_go/internal/api/shared"
	"f_admin_go/internal/api/transactions"
	"f_admin_go/internal/api/users"
	"f_admin_go/internal/config"
	"f_admin_go/internal/db"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.LoadConfig()

	authenticator := shared.NewSimpleAuthenticator(cfg.AuthSecretKey)

	server := &http.Server{
		Addr:         cfg.Port,
		Handler:      routes(authenticator),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

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

func routes(authenticator *shared.SimpleAuthenticator) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/login", func(w http.ResponseWriter, r *http.Request) {
		auth.Login(w, r, authenticator)
	})
	mux.HandleFunc("/api/register", func(w http.ResponseWriter, r *http.Request) {
		auth.Register(w, r, authenticator)
	})

	mux.HandleFunc("/api/assets", shared.HandleEntity(&assets.Handler{}, authenticator))
	mux.HandleFunc("/api/transactions", shared.HandleEntity(&transactions.Handler{}, authenticator))
	mux.HandleFunc("/api/products", shared.HandleEntity(&products.Handler{}, authenticator))
	mux.HandleFunc("/api/users", shared.HandleEntity(&users.Handler{}, authenticator))
	mux.HandleFunc("/api/membership", shared.HandleEntity(&membership.Handler{}, authenticator))
	mux.HandleFunc("/", handleNotFound)

	return mux
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to fantasy tech"))
}
