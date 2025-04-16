package main

import (
	"f_admin_go/internal/api"
	"f_admin_go/internal/db"
	"fmt"
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

	fmt.Println("Server running on port 8080")
	connectStr := ""
	fmt.Println("Connecting to database...")
	if err := db.InitDB(connectStr); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	} else {
		log.Println("Database connection successful")
	}

	// todo: re-connect db if fail

	log.Fatal(server.ListenAndServe())
}

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/asset", api.HandleAssetRequest)
	// mux.HandleFunc("/api/product", api.HandleProductRequest)
	// mux.HandleFunc("/api/transaction", api.HandleTransactionRequest)
	// mux.HandleFunc("/api/user", api.HandleAssetRequest)
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
