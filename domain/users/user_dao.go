package users

import (
	"fmt"

	"github.com/kapilupadhayay/bookstore_users-api/utils/errors"
)

var (
	userdb = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := userdb[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprint("No user exists with ID: %d", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.DateCreated = result.DateCreated
	user.Email = result.Email
	return nil
}

func (user *User) Save() *errors.RestErr {
	if userdb[user.Id] != nil {
		if userdb[user.Id].Email == user.Email {
			return errors.NewAlreadyExistsError(fmt.Sprintf("User ID: %d with email: %s, already exists", user.Id, user.Email))
		}
		userdb[user.Id].Email = user.Email
	}
	userdb[user.Id] = user
	return nil
}
