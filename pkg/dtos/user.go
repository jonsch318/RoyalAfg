package dtos

import (
	"time"

	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
)

// User is the data transfer object of the internal user object
// swagger:model
type User struct {
	// The user id
	// required: true
	ID string `json:"id"`
	// The registration time of the user
	CreatedAt time.Time `json:"created_at"`
	// The time when the user was updated last
	UpdatedAt time.Time `json:"updated_at"`
	// The username of the user
	// required: true
	// min length: 4
	// max length: 100
	Username string `json:"username"`
	// The email of the user
	// required: true
	// min length: 4
	// max length: 100
	// swagger:strfmt email
	Email string `json:"email"`
	// The full name of the user
	// required: true
	// min length: 1
	// max length: 100
	FullName string `json:"fullName"`

	// The unix birthdate of the user
	// required: true
	Birthdate int64 `json:"birthdate"`
}

type UpdateUserDto struct {
	Email string `json:"email"`
	FullName string `json:"fullName"`
}

// NewUserDTO creates a new user dto form the given user
func NewUser(user *models.User) *User {
	return &User{
		ID:        user.ID.Hex(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Birthdate: user.Birthdate,
	}
}