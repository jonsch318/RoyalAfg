package dto

import (
	"github.com/JohnnyS318/RoyalAfgInGo/pkg/models"
	"time"
)

// UserDTO is the data transfer object of the internal user object
// swagger:model
type UserDTO struct {
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

// NewUserDTO creates a new user dto form the given user
func NewUserDTO(user *models.User) *UserDTO {
	return &UserDTO{
		ID:        user.ID.Hex(),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Username:  user.Username,
		Email:     user.Email,
		FullName:  user.FullName,
		Birthdate: user.Birthdate,
	}
}