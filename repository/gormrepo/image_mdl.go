package gormrepo

import (
	"bettersocial/model"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type Image struct {
	Id          *string
	BinaryImage []byte
}

func (p Image) FromModel(data model.Image) *Image {
	return &Image{
		Id:          data.Id,
		BinaryImage: data.BinaryImage,
	}
}

func (p Image) ToModel() *model.Image {
	return &model.Image{
		Id:          p.Id,
		BinaryImage: p.BinaryImage,
	}
}

func (p Image) GetID() *string {
	return p.Id
}

func (p Image) TableName() string {
	return "images"
}

func (p *Image) BeforeCreate(db *gorm.DB) error {
	if p.Id == nil {
		db.Statement.SetColumn("id", ksuid.New().String())
	}
	return nil
}
