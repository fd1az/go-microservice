package domain

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserNotUserFound(t *testing.T) {
	//Initialization:(only if you need)

	//Execution:
	user, err := GetUser(0)

	//Validation:
	assert.Nil(t, user, "we are not expecting a user with id 0")
	assert.NotNil(t, err, "we were expecting an error user id is 0")
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Code)
	assert.EqualValues(t, "user 0 was not found", err.Message)

	//Without testify library
	// if user != nil {
	// 	t.Error("we are not expecting a user with id 0")
	// }
	// if err == nil {
	// 	t.Error("we were expecting an error user id is 0")
	// }
	// if err.StatusCode != http.StatusNotFound {
	// 	t.Error("we were expecting 404 when the user is not found")
	// }
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)
	assert.Nil(t, err, "we are not expecting any error")
	assert.NotNil(t, user, "we were expecting an user")
	assert.EqualValues(t, 123, user.Id)
	assert.EqualValues(t, "Facu", user.FirstName)
	assert.EqualValues(t, "Diaz", user.LastName)
	assert.EqualValues(t, "mail@mail.com", user.Email)
}
