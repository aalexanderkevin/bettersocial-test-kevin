package usecase

import (
	"bettersocial/container"
	"bettersocial/helper"
	"bettersocial/model"

	"bettersocial/repository"
	"context"
)

type Image struct {
	repository.Image
}

func NewImage(c *container.Container) *Image {
	return &Image{
		Image: c.ImageRepo(),
	}
}

func (t *Image) Upload(ctx context.Context, req []byte) (*model.Image, error) {
	logger := helper.GetLogger(ctx).WithField("method", "usecase.Upload")

	res, err := t.Image.Upload(ctx, req)
	if err != nil {
		logger.WithError(err).Warning("Failed insert image")
		return nil, err
	}

	return res, nil
}
