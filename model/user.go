package model

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	Id        *string    `json:"id,omitempty"`
	Username  *string    `json:"username,omitempty"`
	PhotoId   *string    `json:"photo_id,omitempty"`
	Password  *string    `json:"password,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(
		&u,
		validation.Field(&u.Username, validation.Required, is.Alphanumeric),
	)
}
