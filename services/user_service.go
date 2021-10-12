package services

import (
	"github.com/jademnp/go-store-user-api/domain/users"
	"github.com/jademnp/go-store-user-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	err := result.Get()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
