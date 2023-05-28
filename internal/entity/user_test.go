package entity

// Import testify
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user, err := NewUser("John Doe", "johndoe@test.com", "123456")
	assert.Nil(t, err)
	assert.NotEmpty(t, user.ID)
	assert.NotNil(t, user)
	assert.Equal(t, "John Doe", user.Name)
	assert.NotEmpty(t, user.Password)
	assert.Equal(t, "johndoe@test.com", user.Email)
}

func TestUser_ValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "johndoe@test.com", "123456")
	assert.Nil(t, err)
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
	assert.NotEqual(t, user.Password, "123456")
}
