package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("Caique", "caique@gmail.com", "123")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, "Caique", user.Name)
	assert.Equal(t, "caique@gmail.com", user.Email)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Password)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, _ := NewUser("Caique", "caique@gmail.com", "123")
	assert.True(t, user.ValidatePassword("123"))
	assert.False(t, user.ValidatePassword("1234"))
	assert.NotEqual(t, "123", user.Password)
}
