package auth_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken_SignedWithMapClaims(t *testing.T) {
	// Create the claims
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = 1
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix() // Token expires after 1 hour

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	signedToken, err := token.SignedString([]byte(os.Getenv("API_SECRET")))

	// Assertion on the type of error
	if assert.IsType(t, err, nil) {
		fmt.Println("No error")
	}

	// Assertion on the type of signed token
	if assert.IsType(t, signedToken, "signed token should be string") {
		fmt.Println("Signed token is string")
	}
}
