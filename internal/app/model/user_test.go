package model_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zamirka/rest-demo/internal/app/model"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
