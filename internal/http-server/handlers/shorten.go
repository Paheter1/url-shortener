package handlers

import (
	"encoding/json"
	"net/http"
)

type ShortenRequest struct {
	URL string `json:"url"`
}
type ShortenResponse struct {
	Alias string `json:"alias"`
}

func Shorten(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	resp := ShortenResponse{
		Alias: "abc123",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
