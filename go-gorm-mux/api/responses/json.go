package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ResponseJSON is a struct that holds the response data.
type ResponseJSON struct {
	Data interface{} `json:"data"`
	Meta Meta          `json:"meta"`
}

// Meta is a struct that holds the metadata information.
type Meta struct {
	Pagination Pagination `json:"pagination"`
}

// Pagination is a struct that holds pagination information.
type Pagination struct {
	Total  int `json:"total"`
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// JSON sends a JSON response to the client.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR sends a JSON error response to the client.
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
	JSON(w, http.StatusBadRequest, nil)
}
