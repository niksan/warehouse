package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	http.ListenAndServe("localhost:8080", mux)
}

type healthResponse struct {
	Status string `json:"status"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := healthResponse{Status: "ok"}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "encode failed", http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
