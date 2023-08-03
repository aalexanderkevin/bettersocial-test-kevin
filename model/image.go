package model

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Image struct {
	Id          *string `json:"id,omitempty"`
	BinaryImage []byte  `json:"binary_image,omitempty"`
}

func (p Image) Validate() error {
	return validation.ValidateStruct(
		&p,
		validation.Field(&p.BinaryImage, validation.Required),
	)
}
