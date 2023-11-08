package entity_test

import (
	"testing"

	"github.com/ruhancs/hubla-test/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user,err := entity.NewUser("N1")

	assert.Nil(t,err)
	assert.NotNil(t,user)
	assert.Equal(t,user.Name,"N1")
	assert.Equal(t,user.Balance,0.0)
}

func TestInvalidNewUser(t *testing.T) {
	_,err := entity.NewUser("")

	assert.NotNil(t,err)
}