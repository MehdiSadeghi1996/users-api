package users

import (
	"fmt"
	"users-api/utils/date_utils"
	"users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found ", user.Id))
	}

	return nil
}

func (user User) Save() *errors.RestErr {
	user.DateCreated = date_utils.GetNowString()
	return nil
}
