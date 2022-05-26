package users

import (
	"strings"
	"users-api/utils/errors"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.FirstName = strings.TrimSpace(strings.ToLower(user.FirstName))
	if user.FirstName == "" {
		return errors.NewBadRequestError("invalid email FirstName")
	}

	user.LastName = strings.TrimSpace(strings.ToLower(user.LastName))
	if user.LastName == "" {
		return errors.NewBadRequestError("invalid email LastName")
	}

	//user.DateCreated = strings.TrimSpace(strings.ToLower(user.DateCreated))
	//if user.DateCreated == "" {
	//	return errors.NewBadRequestError("invalid email DateCreated")
	//}

	return nil
}
