package controllers

import (
	"encoding/json"
	"net/http"
)

// SendErrorResponse sends an error response with the specified status code and message.
func SendErrorResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}

// SendSuccessResponse sends a success response with the specified status code and message.
func SendSuccessResponse(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"message": message})
}
