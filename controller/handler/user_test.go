package handler_test

import (
	"bettersocial/container"
	"bettersocial/controller/response"
	"bettersocial/helper/test"
	"bettersocial/model"
	"bettersocial/repository/mocks"
	"encoding/json"

	"net/http"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestUser_CheckUsername(t *testing.T) {
	t.Parallel()
	t.Run("ShouldReturn200_WhenNotFoundUsername", func(t *testing.T) {
		t.Parallel()
		// INIT
		username := "username"
		userMock := &mocks.User{}
		userMock.On("GetByUsername", mock.Anything, username).Return(nil, model.NewNotFoundError()).Once()

		router := test.SetupHttpHandler(t, func(appContainer *container.Container) *container.Container {
			appContainer.SetUserRepo(userMock)
			return appContainer
		})

		// CODE UNDER TEST
		w, err := performRequest(router, "GET", "/v1/user/"+username, nil, nil, nil)
		require.NoError(t, err)
		defer printOnFailed(t)(w.Body.String())

		// EXPECTATION
		require.Equal(t, http.StatusOK, w.Code)

		userMock.AssertExpectations(t)
	})

	t.Run("ShouldReturnErrorConflict_WhenGetByUsernameReturnNoError", func(t *testing.T) {
		t.Parallel()
		// INIT
		username := "username"
		userMock := &mocks.User{}
		userMock.On("GetByUsername", mock.Anything, username).Return(&model.User{}, nil).Once()

		router := test.SetupHttpHandler(t, func(appContainer *container.Container) *container.Container {
			appContainer.SetUserRepo(userMock)
			return appContainer
		})

		// CODE UNDER TEST
		w, err := performRequest(router, "GET", "/v1/user/"+username, nil, nil, nil)
		require.NoError(t, err)
		defer printOnFailed(t)(w.Body.String())

		// EXPECTATION
		require.Equal(t, http.StatusConflict, w.Code)

		resBody := response.ErrorResponse{}
		err = json.NewDecoder(w.Body).Decode(&resBody)
		require.NoError(t, err)

		userMock.AssertExpectations(t)
	})

	t.Run("ShouldReturnErrorBadRequest_WhenUsernameIsNotAlphanumeric", func(t *testing.T) {
		t.Parallel()
		// INIT
		username := "username-1"

		router := test.SetupHttpHandler(t, func(appContainer *container.Container) *container.Container {
			return appContainer
		})

		// CODE UNDER TEST
		w, err := performRequest(router, "GET", "/v1/user/"+username, nil, nil, nil)
		require.NoError(t, err)
		defer printOnFailed(t)(w.Body.String())

		// EXPECTATION
		require.Equal(t, http.StatusBadRequest, w.Code)

		resBody := response.ErrorResponse{}
		err = json.NewDecoder(w.Body).Decode(&resBody)
		require.NoError(t, err)

		require.Equal(t, "username should be an alphanumeric", resBody.Message)
	})
}
