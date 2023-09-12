package user

import (
	"context"

	"github.com/jonsch318/royalafg/pkg/models"
)

// SaveUser saves the user to the user service
func (service *User) SaveUser(user *models.User) error {
	message := ToMessage(user)
	_, err := service.Client.SaveUser(context.Background(), message)

	if err != nil {
		return err
	}

	return nil
}
