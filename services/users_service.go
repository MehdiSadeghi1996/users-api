package services

import (
	"users-api/domain/users"
	"users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	createdUser, err := users.User.Save(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
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

func UpdateUser(isPatchOrPartialUpdate bool, user users.User) (*users.User, *errors.RestErr) {

	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPatchOrPartialUpdate {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.FirstName != "" {
			current.LastName = user.LastName
		}
		if user.FirstName != "" {
			current.Email = user.Email
		}

	} else {

		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil

}

func DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{
		Id: userId,
	}
	return user.Delete()
}
