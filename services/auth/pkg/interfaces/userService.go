package interfaces

import "github.com/JohnnyS318/RoyalAfgInGo/pkg/models"

type UserService interface {
	GetUserById(id string) (*models.User, error)
	GetUserByUsernameOrEmail(usernameOrEmail string) (*models.User,error)
	SaveUser(user *models.User) error
}
