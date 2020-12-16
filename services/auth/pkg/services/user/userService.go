package user

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
)

type UserService struct {
	Client protos.UserServiceClient
}

func NewUserService(client protos.UserServiceClient) *UserService{
	return &UserService{
		Client: client,
	}
}



