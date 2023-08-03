package handler

import (
	"bettersocial/container"
	"bettersocial/controller/response"
	"bettersocial/helper"
	"bettersocial/model"
	"bettersocial/usecase"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type User struct {
	appContainer *container.Container
}

func NewUser(appContainer *container.Container) *User {
	return &User{appContainer: appContainer}
}

// @Summary check username
// @Description check username
// @Tags check-username
// @Accept  json
// @Produce  json
// @Param username path string true " "
// @Success 200
// @Failure 401
// @Failure 422
// @Security BearerAuth
// @Router /v1/user/:username [get]
func (u *User) CheckUsername(c *gin.Context) {
	logger := helper.GetLogger(c).WithField("method", "Controller.Handler.CheckUsername")

	// Get username from path
	username := c.Param("username")

	if err := validation.Validate(username, is.Alphanumeric); err != nil {
		response.WriteFailResponse(c, http.StatusBadRequest, errors.New("username should be an alphanumeric"))
		return
	}

	// Action
	userUseCase := usecase.NewUser(u.appContainer)
	err := userUseCase.CheckUsername(c, username)
	if err != nil {
		var e model.Error
		if !errors.As(err, &e) {
			logger.WithError(err).Warning("error get username")
			response.WriteFailResponse(c, http.StatusInternalServerError, err)
		} else {
			response.WriteFailResponse(c, e.Code, e)
		}
		return
	}

	response.WriteSuccessResponse(c, nil)
}
