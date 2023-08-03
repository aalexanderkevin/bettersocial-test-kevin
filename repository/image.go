package repository

import (
	"bettersocial/model"
	"context"
)

type Image interface {
	Upload(context.Context, []byte) (*model.Image, error)
}
