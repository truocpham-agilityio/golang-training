package controllers_test

import (
	"go-gorm-mux/api/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHome(t *testing.T) {
	t.Parallel()

	r, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	controllers.Home(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %v, got %v", http.StatusOK, w.Code)
	}

	expected := "\"Welcome To This Awesome API\"\n"
	if w.Body.String() != expected {
		t.Errorf("Expected body %v, got %v", expected, w.Body.String())
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, expected, w.Body.String())
}
