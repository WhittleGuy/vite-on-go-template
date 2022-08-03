package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/api"
	"server/hiddenFs"
	"time"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	// Load enviroinment variables
	fmt.Printf("[+] Loading environment variables...\n")
	envErr := godotenv.Load("../.env")
	if envErr != nil {
		log.Fatalf("Error loading .env file")
	}

	// Load static assets (your vite app)
	static := os.Getenv("STATIC_ASSETS_DIR")
	fmt.Printf("[+] Retrieving assets from %s\n", static)

	// Create router with chi
	router := chi.NewRouter()

	// Send API routes to their specific router
	router.Mount("/api/", api.NewRouter())

	// Create FileSystem with hiddenFs, obfuscating our directory from users
	fs := http.FileServer(hiddenFs.Dir("../" + static))

	// Send all wildcards to the filesystem
	router.Handle("/*", fs)

	// Create custom server
	srv := &http.Server{
		Handler:      router,
		Addr:         "localhost:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Run server
	fmt.Printf("[+] Starting server...\n")
	srvErr := srv.ListenAndServe()
	if srvErr != nil {
		log.Fatal(srvErr)
	}
}
