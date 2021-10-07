package users

import (
	"fmt"

	"github.com/jademnp/go-store-user-api/utils/errors"
)

var db = make(map[int64]*User)

func (user *User) Save() *errors.RestErr {
	current := db[user.Id]
	if current != nil {
		return errors.BadRequestError(fmt.Sprintf("user id %d already exists", user.Id))
	}
	db[user.Id] = user
	return nil
}
