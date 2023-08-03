package gormrepo

import (
	"bettersocial/model"

	"context"

	"gorm.io/gorm"
)

type ImageRepo struct {
	db *gorm.DB
}

func NewImageRepository(db *gorm.DB) *ImageRepo {
	return &ImageRepo{db}
}

func (u *ImageRepo) Upload(ctx context.Context, image []byte) (*model.Image, error) {
	gormModel := Image{BinaryImage: image}

	if err := u.db.WithContext(ctx).Create(&gormModel).Error; err != nil {
		return nil, err
	}

	return gormModel.ToModel(), nil
}
