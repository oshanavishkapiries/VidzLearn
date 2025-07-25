package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type jsonResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, resp jsonResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

// 200 OK
func Success(w http.ResponseWriter, data interface{}, message string) {
	writeJSON(w, http.StatusOK, jsonResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// 400 Bad Request
func BadRequest(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusBadRequest, jsonResponse{
		Success: false,
		Message: message,
	})
}

// 401 Unauthorized
func Unauthorized(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusUnauthorized, jsonResponse{
		Success: false,
		Message: message,
	})
}

// 403 Forbidden
func Forbidden(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusForbidden, jsonResponse{
		Success: false,
		Message: message,
	})
}

// 404 Not Found
func NotFound(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusNotFound, jsonResponse{
		Success: false,
		Message: message,
	})
}

// 422 Unprocessable Entity
func UnprocessableEntity(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusUnprocessableEntity, jsonResponse{
		Success: false,
		Message: message,
	})
}

// 409 Conflict
func Conflict(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusConflict, jsonResponse{
		Success: false,
		Message: message,
	})
}

// 500 Internal Server Error
func InternalServerError(w http.ResponseWriter, message string) {
	writeJSON(w, http.StatusInternalServerError, jsonResponse{
		Success: false,
		Message: message,
	})
}

func HandleError(w http.ResponseWriter, err error) {
	if apiErr, ok := err.(*APIError); ok {
		writeJSON(w, apiErr.StatusCode, jsonResponse{
			Success: false,
			Message: apiErr.Message,
		})
	} else {
		writeJSON(w, http.StatusInternalServerError, jsonResponse{
			Success: false,
			Message: "Internal server error",
		})
	}
}
