//go:build integration
// +build integration

package gormrepo_test

import (
	"bettersocial/helper/test"
	"bettersocial/model"
	"bettersocial/repository/gormrepo"
	"bettersocial/storage"

	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUserRepository_Add(t *testing.T) {
	t.Run("ShouldInsertUser", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeUser := test.FakeUser(t, nil)

		//-- code under test
		userRepo := gormrepo.NewUserRepository(db)
		addedUser, err := userRepo.Add(context.TODO(), fakeUser)

		//-- assert
		require.NoError(t, err)
		require.NotNil(t, addedUser)
		require.Equal(t, fakeUser.Username, addedUser.Username)
		require.Equal(t, fakeUser.Password, addedUser.Password)
	})

	t.Run("ShouldReturnError_WhenInsertTheSameUsername", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeUser1 := test.FakeUserCreate(t, db, nil)
		fakeUser2 := test.FakeUser(t, func(user model.User) model.User {
			user.Username = fakeUser1.Username
			return user
		})

		//-- code under test
		userRepo := gormrepo.NewUserRepository(db)
		addedUser, err := userRepo.Add(context.TODO(), fakeUser2)

		//-- assert
		require.Error(t, err)
		require.Nil(t, addedUser)
	})
}

func TestUserRepository_GetByUsername(t *testing.T) {
	t.Run("ShouldReturnNotFoundError_WhenTheUsernameIsNotExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		//-- code under test
		userRepo := gormrepo.NewUserRepository(db)
		todo, err := userRepo.GetByUsername(context.TODO(), "invalid-username")
		require.Error(t, err)

		//-- assert
		require.EqualError(t, err, model.NewNotFoundError().Error())
		require.Nil(t, todo)
	})

	t.Run("ShouldGetByUsername_WhenTheUsernameExist", func(t *testing.T) {
		//-- init
		db := storage.PostgresDbConn(&dbName)
		defer cleanDB(t, db)

		fakeUser := test.FakeUserCreate(t, db, nil)

		//-- code under test
		userRepo := gormrepo.NewUserRepository(db)
		todo, err := userRepo.GetByUsername(context.TODO(), *fakeUser.Id)
		require.NoError(t, err)

		//-- assert
		require.NotNil(t, todo)
		require.Equal(t, *fakeUser.Id, *todo.Id)
		require.Equal(t, *fakeUser.Username, *todo.Username)
		require.Equal(t, *fakeUser.Password, *todo.Password)
		require.Equal(t, *fakeUser.PhotoId, *todo.PhotoId)
	})
}
