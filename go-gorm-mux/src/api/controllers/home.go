package controllers

import (
	"go-gorm-mux/src/api/responses"
	"net/http"
)

// Home is the home page of the API.
func Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
