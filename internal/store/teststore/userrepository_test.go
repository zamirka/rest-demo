package teststore_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zamirka/rest-demo/internal/app/model"
	"github.com/zamirka/rest-demo/internal/store"
	"github.com/zamirka/rest-demo/internal/store/teststore"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestFindByEmail(t *testing.T) {
	u := model.TestUser(t)

	s := teststore.New()
	_, err := s.User().FindByEmail(u.Email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
