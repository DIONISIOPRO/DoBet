package domain

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	user := User{
		Phone_number: "23546",
	}
	err := user.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "the lenght of number should be 9", err.Error())
	user.Phone_number = "214d57854"
	err = user.Validate()
	assert.NotNil(t, err)
	assert.Equal(t, "your number is not valid", err.Error())
	user.Phone_number = "123456789"
	err = user.Validate()
	assert.NoError(t, err)
}

func TestAddRefeshToken(t *testing.T) {
	user := User{
		Phone_number: "23546",
	}
	assert.Equal(t, 0, len(user.RefreshTokens))
	user.AddRefreshToken("token")
	assert.Equal(t, 1, len(user.RefreshTokens))
	user.AddRefreshToken("token")
	user.AddRefreshToken("token")
	user.AddRefreshToken("token")
	user.AddRefreshToken("token")
	user.AddRefreshToken("token")
	user.AddRefreshToken("token")
	assert.Equal(t, 7, len(user.RefreshTokens))
}

func TestPromoteToAdmin(t *testing.T) {
	user := User{
		Phone_number: "123456789",
	}
	assert.Equal(t, false, user.IsAdmin)
	err := user.PromoteToAdmin()
	assert.NoError(t, err)
	assert.Equal(t, true, user.IsAdmin)
}

func TestUpdate(t *testing.T) {
	user := User{
		Phone_number: "123456789",
	}
	time := user.Updated_at
	err := user.Update()
	time1 := user.Updated_at
	assert.NoError(t, err)
	assert.NotEqual(t, time, time1)
}
