package services

import (
	"github.com/jademnp/go-store-user-api/domain/users"
	"github.com/jademnp/go-store-user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
