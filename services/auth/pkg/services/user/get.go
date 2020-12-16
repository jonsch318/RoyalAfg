package user

import (
	"context"
	"errors"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (service *UserService)GetUserByUsernameOrEmail(usernameOrEmail string) (*models.User, error) {
	if usernameOrEmail == ""{
		return nil, errors.New("The given username or email is empty")
	}

	var message *protos.User
	isEmail := validation.Validate(usernameOrEmail, is.EmailFormat) == nil
	if isEmail {
		//TBA
	}else {
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

func (service *UserService)GetUserById(id string) (*models.User,error){
	message, err := service.Client.GetUserById(context.Background(), &protos.GetUser{
		ApiKey: "",
		Identifier: id,
	})
	if err != nil {
		return nil, err
	}
	user := FromMessage(message)
	return user, nil
}