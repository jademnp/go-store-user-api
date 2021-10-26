package users

import (
	"github.com/jademnp/go-store-user-api/datasources/postgresql/users_db"
	"github.com/jademnp/go-store-user-api/utils/date_utils"
	"github.com/jademnp/go-store-user-api/utils/errors"
	"github.com/jademnp/go-store-user-api/utils/postgressql_utils"
)

const (
	queryInsertUser = "insert into users (first_name,last_name,email,created_date) values ($1,$2,$3,$4) RETURNING id;"
	querySelectUser = "select first_name,last_name,email,created_date from users where id = $1"
)

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return postgressql_utils.ParseError(err)
	}
	defer stmt.Close()
	user.CreatedDate = date_utils.GetNowString()
	insertResult := stmt.QueryRow(user.FirstName, user.LastName, user.Email, user.CreatedDate)
	var userId int64
	err = insertResult.Scan(&userId)
	if err != nil {
		return postgressql_utils.ParseError(err)
	}
	user.Id = userId
	return nil
}

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(querySelectUser)
	if err != nil {
		return postgressql_utils.ParseError(err)
	}
	defer stmt.Close()
	result := stmt.QueryRow(user.Id)
	err = result.Scan(&user.FirstName, &user.LastName, &user.Email, &user.CreatedDate)
	if err != nil {
		return postgressql_utils.ParseError(err)
	}
	return nil
}
