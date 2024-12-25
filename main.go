package main

import (
	"encoding/json"
	"fmt"
	"main/util"

	http "github.com/useflyent/fhttp"
)

func main() {
	http.HandleFunc("/api/detect", detectHandler)
	fmt.Println("Server is running on port 5000")
	http.ListenAndServe(":5000", nil)
}

func detectHandler(w http.ResponseWriter, r *http.Request) {
	// Handle CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	var payload struct {
		URL string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	if payload.URL == "" {
		http.Error(w, "URL parameter is missing", http.StatusBadRequest)
		return
	}

	services, err := util.ParseWebsite(payload.URL)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing website: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(services)
} 