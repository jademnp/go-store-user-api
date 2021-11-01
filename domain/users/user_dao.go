package users

import (
	"github.com/jademnp/go-store-user-api/datasources/postgresql/users_db"
	"github.com/jademnp/go-store-user-api/utils/date_utils"
	"github.com/jademnp/go-store-user-api/utils/errors"
	"github.com/jademnp/go-store-user-api/utils/postgressql_utils"
)

const (
	indexUniqueEmail = "users_un"
	queryInsertUser  = "insert into users (first_name,last_name,email,created_date) values ($1,$2,$3,$4) RETURNING id;"
	querySelectUser  = "select first_name,last_name,email,created_date from users where id = $1"
	queryUpdateUser  = "UPDATE users SET first_name=$1, last_name=$2, email=$3 WHERE id=$4;"
	queryDeleteUser  = "DELETE FROM users WHERE id=$1;"
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
func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.InternalServerError("error when tying to update user")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.InternalServerError("error when tying to update user2")
	}
	return nil
}

func (user *User) Delete() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.InternalServerError("error when tying to update user")
	}
	defer stmt.Close()

	if _, err = stmt.Exec(user.Id); err != nil {
		return errors.InternalServerError("error when tying to save user")
	}
	return nil
}
