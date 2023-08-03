package gormrepo

import (
	"time"

	"bettersocial/model"

	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)

type User struct {
	Id        *string
	Username  *string
	ImageId   *string
	Password  *string
	CreatedAt *time.Time
}

func (u User) FromModel(data model.User) *User {
	return &User{
		Id:       data.Id,
		Username: data.Username,
		ImageId:  data.ImageId,
		Password: data.Password,
	}
}

func (u User) ToModel() *model.User {
	return &model.User{
		Id:        u.Id,
		Username:  u.Username,
		ImageId:   u.ImageId,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
	}
}

func (u User) GetID() *string {
	return u.Id
}

func (u User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	if u.Id == nil {
		db.Statement.SetColumn("id", ksuid.New().String())
	}
	return nil
}
