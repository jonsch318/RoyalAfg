package user

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/log"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/protos"
	"github.com/JohnnyS318/RoyalAfgInGo/services/auth/pkg/serviceconfig"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

//IUser is responsible to communicate with the user service
type IUser interface {
	//GetUserById returns the user with the given id if found
	GetUserById(id string) (*models.User, error)
	//GetUserById returns the user with the given username or email if found
	GetUserByUsernameOrEmail(usernameOrEmail string) (*models.User, error)
	//SaveUser saves the new user to the user service
	SaveUser(user *models.User) error
}

//User communicates with the user services
type User struct {
	Client protos.UserServiceClient
	conn   *grpc.ClientConn
}

//NewUser creates a new user service
func NewUser() (*User, error) {

	log.Logger.Infof("Auth service url %v trying to connect", viper.GetString(serviceconfig.UserServiceUrl))
	conn, err := grpc.Dial(viper.GetString(serviceconfig.UserServiceUrl), grpc.WithInsecure())

	if err != nil {
		return nil, err
	}
	state := conn.GetState()
	log.Logger.Infow("Calling state", "state", state.String())

	client := protos.NewUserServiceClient(conn)

	return &User{
		Client: client,
		conn:   conn,
	}, nil
}

func (u *User) Close() {
	u.conn.Close()
}
