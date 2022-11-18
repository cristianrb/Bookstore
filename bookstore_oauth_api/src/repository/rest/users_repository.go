package rest

import (
	"cristianrb/src/domain/users"
	"cristianrb/src/utils/errors"
	"github.com/go-resty/resty/v2"
)

var (
	usersRestClient = resty.New()
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type restUsersRepository struct {
}

func NewRepository() RestUsersRepository {
	return &restUsersRepository{}
}

func (r *restUsersRepository) LoginUser(email string, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	var user users.User
	var restErr errors.RestErr
	response, err := usersRestClient.R().
		SetBody(request).
		SetResult(&user).
		SetError(&restErr).
		Post("http://localhost:8080/users/login")

	if response == nil || response.RawResponse == nil || err != nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user")
	}

	return &user, nil
}
