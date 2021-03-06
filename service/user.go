package service

import (
	"errors"

	"github.com/izzanzahrial/library-service/entity"
)

var userID int

type user struct {
	users map[int]*entity.User
}

func (u *user) RegisterUser(name string) error {
	userID += 1

	u.users[userID] = &entity.User{
		ID:   userID,
		Name: name,
	}

	return nil
}

func (u *user) GetUser(userID int) (*entity.User, error) {
	user, ok := u.users[userID]
	if !ok {
		return &entity.User{}, errors.New("cannot find user")
	}

	return user, nil
}
