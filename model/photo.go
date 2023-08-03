package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Photo struct {
	Id     *string `json:"id,omitempty"`
	Binary *string `json:"binary,omitempty"`
}

func (p Photo) Validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.Binary, validation.Required),
	)
}
