package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

// JSON representation of a customized message
type HelloResponse struct {
	Message string `json:"message"`
}

// Creates a JSON response
func jsonResponse(w http.ResponseWriter, data interface{}, c int) {
	dj, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(c)
	fmt.Fprintf(w, "%s", dj)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	response := HelloResponse{
		Message: "Hello, World!",
	}
	jsonResponse(w, response, http.StatusOK)
}

// returns a personalized JSON message
func HelloName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	response := HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", name),
	}
	jsonResponse(w, response, http.StatusOK)
}

// NewRouter returns an HTTP handler that implements routes for the API
func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Register API routes
	r.Get("/", helloWorld)
	r.Get("/{name}", HelloName)
	
	return r
}