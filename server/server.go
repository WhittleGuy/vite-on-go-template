package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"server/api"
	"server/hiddenFs"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Printf("[+] Loading environment variables...\n")
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	
	static := os.Getenv("STATIC_ASSETS_DIR")
	fmt.Printf("[+] Retrieving assets from %s\n", static)

	router := chi.NewRouter()
	router.Mount("/api/", api.NewRouter())
	fs := http.FileServer(hiddenFs.Dir("../"+static))
	router.Handle("/*", fs)
	fmt.Printf("[+] Starting server...\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}