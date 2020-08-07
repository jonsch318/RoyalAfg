package dtos

import (
	"encoding/json"
	"io"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type LoginUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (dto *LoginUser) FromJSON(r io.Reader) error {
	decoder := json.NewDecoder(r)
	return decoder.Decode(dto)
}

func (dto *LoginUser) ToJSON(w io.Writer) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(dto)
}

func (dto LoginUser) Validate() error {
	return validation.ValidateStruct(&dto,
		validation.Field(&dto.Username, validation.Required, validation.Length(4, 100)),
		validation.Field(&dto.Password, validation.Required, validation.Length(4, 100)),
	)
}
