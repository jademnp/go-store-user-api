package users

import (
	"fmt"
	"strings"

	"github.com/jademnp/go-store-user-api/datasources/postgresql/users_db"
	"github.com/jademnp/go-store-user-api/utils/date_utils"
	"github.com/jademnp/go-store-user-api/utils/errors"
)

var db = make(map[int64]*User)

const (
	indexUniqueEmail = "users_un"
	quereyInsertUser = "insert into users (first_name,last_name,email,created_date) values ($1,$2,$3,$4) RETURNING id;"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(quereyInsertUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	user.CreatedDate = date_utils.GetNowString()
	insertResult := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.CreatedDate)
	var userId int64
	err = insertResult.Scan(&userId)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.InternalServerError(fmt.Sprintf("email %s is already exist", user.Email))
		}
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	if err != nil {
		return errors.InternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}
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
