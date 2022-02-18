package services

import (
	"net/http"
	"users-api/domain/users"
	"users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	if userId <= 0 {
		return nil, errors.NewBadRequestError("invalid user id")
	}
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil

}
