package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var store = make(map[string]bool)
var loginManager = NewLogInStateManager()

func TestIsLogIn(t *testing.T) {
	loginManager.stateStore = store
	loginManager.LogIn("123")
	login := loginManager.IsLogIn("123")
	assert.Equal(t, login, true)
	loginManager.Logout("123")
	assert.Equal(t, login, true)
	login = loginManager.IsLogIn("123")
	assert.Equal(t, login, false)
}
