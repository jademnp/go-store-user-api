package users

import (
	"fmt"
	"strings"

	"github.com/jademnp/go-store-user-api/datasources/postgresql/users_db"
	"github.com/jademnp/go-store-user-api/utils/date_utils"
	"github.com/jademnp/go-store-user-api/utils/errors"
)

const (
	indexUniqueEmail = "users_un"
	queryInsertUser  = "insert into users (first_name,last_name,email,created_date) values ($1,$2,$3,$4) RETURNING id;"
	querySelectUser  = "select first_name,last_name,email,created_date from users where id = $1"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
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
	stmt, err := users_db.Client.Prepare(querySelectUser)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	err = result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.CreatedDate)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	return nil
}
