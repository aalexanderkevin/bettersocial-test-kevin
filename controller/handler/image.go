package handler

import (
	"bettersocial/container"
	"bettersocial/controller/response"
	"bettersocial/helper"
	"bettersocial/model"
	"bettersocial/usecase"
	"encoding/base64"

	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Image struct {
	appContainer *container.Container
}

func NewImage(appContainer *container.Container) *Image {
	return &Image{appContainer: appContainer}
}

// @Summary upload
// @Description upload
// @Tags upload
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 401
// @Failure 422
// @Security BearerAuth
// @Router /v1/upload [post]
func (u *Image) Upload(c *gin.Context) {
	logger := helper.GetLogger(c).WithField("method", "Controller.Handler.Upload")

	// Decode the request body into binary data
	imgData, err := decodeBase64Image(c)
	if err != nil {
		logger.WithError(err).Warning("invalid image data")
		response.WriteFailResponse(c, http.StatusBadRequest, errors.New("Invalid image data"))
		return
	}

	// Action
	imageUseCase := usecase.NewImage(u.appContainer)
	res, err := imageUseCase.Upload(c, imgData)
	if err != nil {
		var e model.Error
		if !errors.As(err, &e) {
			logger.WithError(err).Warning("error upload")
			response.WriteFailResponse(c, http.StatusInternalServerError, err)
		} else {
			response.WriteFailResponse(c, e.Code, e)
		}
		return
	}

	response.WriteSuccessResponse(c, res)
}

func decodeBase64Image(c *gin.Context) ([]byte, error) {
	// Get the binary data from the request body
	binaryData := make([]byte, c.Request.ContentLength)
	_, err := c.Request.Body.Read(binaryData)
	if err != nil {
		return nil, err
	}

	// Decode the base64 data
	decodedData := make([]byte, base64.StdEncoding.DecodedLen(len(binaryData)))
	_, err = base64.StdEncoding.Decode(decodedData, binaryData)
	if err != nil {
		return nil, err
	}

	return decodedData, nil
}
