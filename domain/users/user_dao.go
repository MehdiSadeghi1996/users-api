package users

import (
	"fmt"
	"users-api/datasources/mysql/users_db"
	"users-api/utils/date_utils"
	"users-api/utils/errors"
)

const (
	queryInsertUser = "INSERT INTO users(first_name,last_name,email,date_created) VALUES (?,?,?,?);"
	queryGetUser    = "SELECT id,first_name,last_name,email,date_created FROM users WHERE id=?;"
	queryUpdateUser = "UPDATE users SET firs_name=?,last_name=?,email=? WHERE id = ?;"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {

	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); err != nil {
		fmt.Println(err)
		return errors.NewInternalError(fmt.Sprintf("error when try to get user %d", user.Id))
	}
	return nil

}

func (user User) Save() (*User, *errors.RestErr) {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()
	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}
	userId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, errors.NewInternalError(err.Error())
	}
	user.Id = userId

	return &user, nil
}

func (user User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(&user.FirstName, &user.LastName, &user.Email, &user.Id)
	if err != nil {
		return errors.NewInternalError(err.Error())
	}

	return nil
}
