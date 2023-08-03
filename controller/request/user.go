package request

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	ImageId  *string `json:"image,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Username, validation.Required, is.Alphanumeric),
		validation.Field(&u.Password, validation.Required),
	)
}
