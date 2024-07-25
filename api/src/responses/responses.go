package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// Request_JSON returns a response writer with JSON defined status code and message
func RequestJSON(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if message != nil {
		if err := json.NewEncoder(w).Encode(message); err != nil {
			log.Fatal(err)
		}
	}
}

// ErrorMessage format a message to response when receive an error execution
func ErrorMessage(w http.ResponseWriter, statusCode int, error_message error) {
	RequestJSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: error_message.Error(),
	})
}

// SuccessMessage format a message to response when a success execution
func SuccessMessage(w http.ResponseWriter, statusCode int, success_message string) {
	RequestJSON(w, statusCode, struct {
		Success string `json:"success"`
	}{
		Success: success_message,
	})
}
