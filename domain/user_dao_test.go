package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNoUserFound(t *testing.T) {
	user, err := GetUser(0)

	assert.Nil(t, user, "we were not expecting a user with id 0")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 does not exist", err.Message)
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	assert.NotNil(t, user)
	assert.Nil(t, err)

	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Adrian", user.FirstName)
	assert.EqualValues(t, "Pod", user.LastName)
	assert.EqualValues(t, "apo@wp.pl", user.Email)
}
