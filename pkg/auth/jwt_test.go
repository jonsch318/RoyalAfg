package auth

import (
	"log"
	"testing"
	"time"

	"github.com/jonsch318/royalafg/pkg/config"
	"github.com/jonsch318/royalafg/pkg/models"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SetDefaults() {
	viper.SetDefault(config.SessionCookieName, "Session")
	viper.SetDefault(config.JWTSigningKey, "TestSecret")
	viper.SetDefault(config.JWTIssuer, "Royalafg")
	viper.SetDefault(config.JWTExpiresAt, time.Hour*24*7)
}

func TestGetJwt(t *testing.T) {
	SetDefaults()
	user := models.NewUser("test", "test@test.com", "test test", time.Now().Unix())
	user.ID, _ = primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
	t.Run("Should_Succeed", func(t *testing.T) {

		jwt, err := GetJwt(user)
		assert.Nil(t, err)
		assert.NotEmpty(t, jwt)

		assert.Nil(t, CheckSignature(jwt))
	})
}

func TestGenerateCookie(t *testing.T) {
	token := "abc"
	cookie := GenerateCookie(token, true)
	log.Printf("%v", cookie)
	assert.Equal(t, token, cookie.Value)
	assert.Equal(t, "/", cookie.Path)
	assert.True(t, cookie.HttpOnly)
	assert.Equal(t, "Session", cookie.Domain)
}

func TestWrapToken(t *testing.T) {
	token := "abc"
	wrap := WrapToken(token)
	assert.Equal(t, "Bearer abc", wrap)
}
