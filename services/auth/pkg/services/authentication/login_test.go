package authentication

import (
	"testing"
	"time"

	"github.com/Kamva/mgm"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	mocks "github.com/JohnnyS318/RoyalAfgInGo/mocks/services/auth/pkg/interfaces"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/config"
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
)

func TestLogin(t *testing.T) {
	viper.SetDefault(config.JWTExpiresAt, time.Minute * 60)
	viper.SetDefault(config.JWTIssuer, "example.com")
	viper.SetDefault(config.JWTSigningKey, "testkey")

	mockUserService := &mocks.UserService{}
	auth := NewService(mockUserService)
	user := &models.User{
		DefaultModel: mgm.DefaultModel{},
		Username:     "",
		Email:        "test",
		Hash:         "test",
		FullName:     "test test",
	}

	t.Run("TestCredentialsValid", func(t *testing.T) {
		mockUserService.On("", ).Return()
		user, token, err := auth.Login("testUser", "testPassword")

		assert.Nil(t, err)
		assert.

	})
}