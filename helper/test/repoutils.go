package test

import (
	"bettersocial/helper"
	"bettersocial/model"
	"bettersocial/repository/gormrepo"

	"context"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func FakeUser(t *testing.T, cb func(user model.User) model.User) model.User {
	t.Helper()

	fakeRp := model.User{
		Username: helper.Pointer(fake.CharactersN(10)),
		PhotoId:  helper.Pointer(fake.CharactersN(10)),
		Password: helper.Pointer(fake.CharactersN(10)),
	}
	if cb != nil {
		fakeRp = cb(fakeRp)
	}
	return fakeRp
}

func FakeUserCreate(t *testing.T, db *gorm.DB, callback func(user model.User) model.User) *model.User {
	t.Helper()

	fakeData := FakeUser(t, callback)

	repo := gormrepo.NewUserRepository(db)
	user, err := repo.Add(context.TODO(), fakeData)
	require.NoError(t, err)

	return user
}
