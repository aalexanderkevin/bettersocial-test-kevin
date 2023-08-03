package usecase

import (
	"bettersocial/container"
	"bettersocial/helper"
	"bettersocial/model"

	"bettersocial/repository"
	"context"
)

type User struct {
	repository.User
}

func NewUser(c *container.Container) *User {
	return &User{
		User: c.UserRepo(),
	}
}

func (u *User) CheckUsername(ctx context.Context, username string) error {
	logger := helper.GetLogger(ctx).WithField("method", "usecase.CheckUsername")

	_, err := u.User.GetByUsername(ctx, username)
	if err != nil {
		// return nil if username not exist
		if model.IsNotFoundError(err) {
			return nil
		}
		logger.WithError(err).Warning("Failed getById todo")
		return err
	}

	// return error if username exist
	return model.NewDuplicateError()
}

