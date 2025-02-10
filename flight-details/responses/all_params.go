package responses

import (
	"encoding/json"
	"log"
)

// FormatSuccessResponse formats the response as JSON
func FormatSuccessResponse(data string) map[string]interface{} {
	var response map[string]interface{}
	err := json.Unmarshal([]byte(data), &response)
	if err != nil {
		log.Println("Error formatting response:", err)
		return map[string]interface{}{
			"status":  "error",
			"message": "Response parsing failed",
		}
	}

	return map[string]interface{}{
		"status": "success",
		"data":   response,
	}
}

// SuccessResponse defines the structure for a successful response
type SuccessResponse struct {
	Status string      `json:"status"` // e.g., "success"
	Data   interface{} `json:"data"`   // Actual flight search result
}

// ErrorResponse defines the structure for an error response
type ErrorResponse struct {
	Status  string `json:"status"`  // e.g., "error"
	Message string `json:"message"` // Error details
}

// FormatErrorResponse formats an error response
func FormatErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		Status:  "error",
		Message: message,
	}
}
