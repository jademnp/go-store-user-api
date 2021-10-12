package users

import (
	"fmt"

	"github.com/jademnp/go-store-user-api/utils/date_utils"
	"github.com/jademnp/go-store-user-api/utils/errors"
)

var db = make(map[int64]*User)

func (user *User) Save() *errors.RestErr {
	current := db[user.Id]
	if current != nil {
		return errors.BadRequestError(fmt.Sprintf("user id %d already exists", user.Id))
	}
	user.CreatedDate = date_utils.GetNowString()
	db[user.Id] = user
	return nil
}

func (user *User) Get() *errors.RestErr {
	result := db[user.Id]
	if result == nil {
		return errors.NotFoundError(fmt.Sprintf("user id %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.CreatedDate = result.CreatedDate
	return nil
}
