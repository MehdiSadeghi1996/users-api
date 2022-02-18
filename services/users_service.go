package services

import (
	"net/http"
	"users-api/domain/users"
	"users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return nil, &errors.RestErr{
		Status: http.StatusInternalServerError,
	}
}
