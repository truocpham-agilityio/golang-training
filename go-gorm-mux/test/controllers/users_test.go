package controllers_test

import (
	"bytes"
	"encoding/json"
	"go-gorm-mux/api/controllers"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_WithFailCreateUser(t *testing.T) {
	t.Parallel()

	invalidCreateUser, err := json.Marshal(map[string]string{
		"name":     "",
		"email":    "",
		"password": "",
	})

	if err != nil {
		t.Errorf("Invalid create user: %v", err)
	}

	r, _ := http.NewRequest("POST", "/users", nil)
	r.Header.Set("Content-Type", "application/json")
	r.Body = ioutil.NopCloser(bytes.NewBuffer(invalidCreateUser))
	defer r.Body.Close()

	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")

	controllers.CreateUser(w, r)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}
