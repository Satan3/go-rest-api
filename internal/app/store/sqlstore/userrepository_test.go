package sqlstore_test

import (
	"github.com/Satan3/go-rest-api/internal/app/model"
	store "github.com/Satan3/go-rest-api/internal/app/store/errors"
	"github.com/Satan3/go-rest-api/internal/app/store/sqlstore"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := model.TestUser()
	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	email := "test@mail.ru"
	s := sqlstore.New(db)

	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	u := model.TestUser()
	u.Email = email
	u, err = s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
