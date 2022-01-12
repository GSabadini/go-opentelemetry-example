package domain

import (
	"context"
)

type UserCreator interface {
	Create(context.Context, User) error
}

type User struct {
	id string
}

func NewUser(id string) User {
	return User{id: id}
}

func (u User) ID() string {
	return u.id
}
