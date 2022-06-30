package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go-gorm-mux/api/auth"
	"go-gorm-mux/api/database"
	"go-gorm-mux/api/models"
	"go-gorm-mux/api/responses"
	"go-gorm-mux/api/utils/formaterror"

	"golang.org/x/crypto/bcrypt"
)

// Login controller used to login a user.
func Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

// SignIn represents the login of a user.
func SignIn(email, password string) (string, error) {
	var err error

	user := models.User{}

	err = database.DB.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
