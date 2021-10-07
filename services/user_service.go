package services

import (
	"net/http"

	"github.com/jademnp/go-store-user-api/domain/users"
	"github.com/jademnp/go-store-user-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return nil, &errors.RestErr{Message: "cannot create", Status: http.StatusInternalServerError, Error: "cant_create"}
}
