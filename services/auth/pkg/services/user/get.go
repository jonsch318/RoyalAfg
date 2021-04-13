package user

import (
	"context"
	"errors"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

//GetUserByUsernameOrEmail checks whether the input is a email or a username and uses the functions for either one of them
func (service *User) GetUserByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	if usernameOrEmail == "" {
		return nil, errors.New("The given username or email is empty")
	}

	var message *protos.User
	isEmail := validation.Validate(usernameOrEmail, is.EmailFormat) == nil
	if isEmail {
		//TBA
	} else {
		var err error
		message, err = service.Client.GetUserByUsername(context.Background(), &protos.GetUser{
			ApiKey:     "",
			Identifier: usernameOrEmail,
		})
		if err != nil {
			return nil, err
		}
	}

	user := FromMessage(message)
	return user, nil
}

//GetUserById finds the user with the given id via the user service
func (service *User) GetUserById(id string) (*models.User, error) {
	message, err := service.Client.GetUserById(context.Background(), &protos.GetUser{
		ApiKey:     "",
		Identifier: id,
	})
	if err != nil {
		return nil, err
	}
	user := FromMessage(message)
	return user, nil
}
