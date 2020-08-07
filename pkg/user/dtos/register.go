package dtos

import (
	"encoding/json"
	"io"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// RegisterUser dto
type RegisterUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// FromJSON reads from the given io reader and decodes it if possible to a RegisterUser dto, else it returns an error.
func (dto *RegisterUser) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(dto)
}

// ToJSON serializes the RegisterUser dto and writes it to the given io writer.
func (dto *RegisterUser) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(dto)
}

// Validate validates if the RegisterUser dto matches all the user requirements
func (dto RegisterUser) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Email, is.Email),
	)
}
