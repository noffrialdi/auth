package custerr

import (
	"encoding/json"
	"net/http"
)

// Struktur respons sukses
type SuccessResponse struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Struktur respons error
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// Fungsi untuk mengirim respons sukses
func RespondWithSuccess(w http.ResponseWriter, message string, data interface{}) {
	response := SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Fungsi untuk mengirim respons error
func RespondWithError(w http.ResponseWriter, code int, message string) {
	response := ErrorResponse{
		Status:  "error",
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}
