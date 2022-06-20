package service

import (
	"errors"

	"github.com/izzanzahrial/library-service/entity"
)

var userID int

type user struct {
	users map[string]*entity.User
}

func (u *user) RegisterUser(name, role string) error {
	userID += 1

	u.users[name] = &entity.User{
		ID:   userID,
		Name: name,
	}

	return nil
}

func (u *user) GetUser(name string) (*entity.User, error) {
	user, ok := u.users[name]
	if !ok {
		return &entity.User{}, errors.New("cannot find user")
	}

	return user, nil
}
