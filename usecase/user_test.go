package usecase_test

import (
	"bettersocial/container"
	"bettersocial/model"
	"bettersocial/repository/mocks"
	"bettersocial/usecase"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUser_CheckUsername(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturnError_WhenErrorGetUsername", func(t *testing.T) {
		t.Parallel()
		// INIT
		username := "username"

		userMock := &mocks.User{}
		userMock.On("GetByUsername", mock.Anything, username).Return(nil, errors.New("error get username")).Once()

		appContainer := container.Container{}
		appContainer.SetTodoRepo(userMock)

		// CODE UNDER TEST
		uc := usecase.NewUser(&appContainer)
		err := uc.CheckUsername(context.Background(), username)
		require.Error(t, err)
		require.EqualError(t, err, "error get username")

		userMock.AssertExpectations(t)
	})

	t.Run("ShouldReturnErrorDuplicate_WhenGetUsernameGotNoError", func(t *testing.T) {
		t.Parallel()
		// INIT
		username := "username"

		userMock := &mocks.User{}
		userMock.On("GetByUsername", mock.Anything, username).Return(&model.User{}, nil).Once()

		appContainer := container.Container{}
		appContainer.SetTodoRepo(userMock)

		// CODE UNDER TEST
		uc := usecase.NewUser(&appContainer)
		err := uc.CheckUsername(context.Background(), username)
		require.Error(t, err)
		require.True(t, model.IsDuplicateError(err))

		userMock.AssertExpectations(t)
	})

	t.Run("ShouldReturnNil_WhenGetUsernameGotErrorNotFound", func(t *testing.T) {
		t.Parallel()
		// INIT
		username := "username"

		userMock := &mocks.User{}
		userMock.On("GetByUsername", mock.Anything, username).Return(nil, model.NewNotFoundError()).Once()

		appContainer := container.Container{}
		appContainer.SetTodoRepo(userMock)

		// CODE UNDER TEST
		uc := usecase.NewUser(&appContainer)
		err := uc.CheckUsername(context.Background(), username)
		require.NoError(t, err)

		userMock.AssertExpectations(t)
	})
}

