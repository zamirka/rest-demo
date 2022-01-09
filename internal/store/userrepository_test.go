package store_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zamirka/rest-demo/internal/app/model"
	"github.com/zamirka/rest-demo/internal/store"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestFindByEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u := model.TestUser(t)

	_, err := s.User().FindByEmail(u.Email)
	assert.Error(t, err)

	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
