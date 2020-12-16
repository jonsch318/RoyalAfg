package user

import (
	"context"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
)

func (service *UserService)SaveUser(user *models.User) error{
	message := ToMessage(user)
	_, err := service.client.SaveUser(context.Background(), message)

	if err != nil {
		return err
	}

	return nil
}