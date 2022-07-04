package formaterror_test

import (
	"go-gorm-mux/api/utils/formaterror"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatError_AlwaysReturnError(t *testing.T) {
	expectation := formaterror.FormatError("")

	assert.Error(t, expectation, "should always return error")
}

func TestFormatError_WithNameAlreadyTaken(t *testing.T) {
	expectation := formaterror.FormatError("name already taken")

	assert.EqualErrorf(t, expectation, "name already taken", "should return error with name already taken")
}

func TestFormatError_WithEmailAlreadyTaken(t *testing.T) {
	expectation := formaterror.FormatError("email already taken")

	assert.EqualError(t, expectation, "email already taken", "should return error with email already taken")
}

func TestFormatError_WithTitleAlreadyTaken(t *testing.T) {
	expectation := formaterror.FormatError("title already taken")

	assert.EqualError(t, expectation, "title already taken", "should return error with title already taken")
}

func TestFormatError_WithIncorrectPassword(t *testing.T) {
	expectation := formaterror.FormatError("hashedPassword")

	assert.EqualError(t, expectation, "incorrect password", "should return error with incorrect password")	
}

func TestFormatError_WithIncorrectDetails(t *testing.T) {
	expectation := formaterror.FormatError("incorrect details")

	assert.EqualError(t, expectation, "incorrect details", "should return error with incorrect details")
}
