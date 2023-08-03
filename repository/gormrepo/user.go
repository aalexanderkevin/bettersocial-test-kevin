package gormrepo

import (
	"bettersocial/model"

	"context"
	"errors"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{db}
}

func (u *UserRepo) Add(ctx context.Context, user model.User) (*model.User, error) {
	gormModel := User{}.FromModel(user)

	if err := u.db.WithContext(ctx).Create(&gormModel).Error; err != nil {
		return nil, err
	}

	return gormModel.ToModel(), nil
}

func (u UserRepo) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	user := User{
		Username: &username,
	}

	err := u.db.WithContext(ctx).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, model.NewNotFoundError()
		}
		return nil, err
	}

	return user.ToModel(), nil
}
