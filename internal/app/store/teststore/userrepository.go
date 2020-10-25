package teststore

import (
	"github.com/Satan3/go-rest-api/internal/app/model"
	store "github.com/Satan3/go-rest-api/internal/app/store/errors"
)

type UserRepository struct {
	store *Store
	users map[string]*model.User
}

func (r *UserRepository) Create(u *model.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.BeforeCreate(); err != nil {
		return err
	}

	r.users[u.Email] = u
	u.ID = len(r.users)

	return nil
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	if user, found := r.users[email]; found {
		return user, nil
	}
	return nil, store.ErrRecordNotFound
}
