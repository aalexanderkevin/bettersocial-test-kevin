package repository

import (
	"bettersocial/model"

	"context"
)

type User interface {
	Add(context.Context, model.User) (*model.User, error)
	GetByUsername(context.Context, string) (*model.User, error)
}
