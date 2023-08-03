package controller

import (
	"bettersocial/config"
	"bettersocial/container"
	"bettersocial/controller/handler"
	"bettersocial/controller/middleware"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type HttpServer interface {
	Start() error
	GetHandler() (http.Handler, error)
}

type httpServer struct {
	config      config.Config
	engine      *gin.Engine
	controllers controllers
}

type controllers struct {
	user  handler.User
	image handler.Image
}

func NewHttpServer(container *container.Container) *httpServer {
	gin.SetMode(gin.ReleaseMode)
	if strings.ToLower(container.Config().LogLevel) == gin.DebugMode {
		gin.SetMode(gin.DebugMode)
	}

	engine := newGinEngine()

	controllers := controllers{
		*handler.NewUser(container),
		*handler.NewImage(container),
	}
	requestHandler := &httpServer{container.Config(), engine, controllers}
	requestHandler.setupRouting()

	return requestHandler
}

func (h *httpServer) Start() error {
	return h.engine.Run(fmt.Sprintf("%s:%s", h.config.Service.Host, h.config.Service.Port))
}

func (h *httpServer) GetHandler() (http.Handler, error) {
	return h.engine, nil
}

func newGinEngine() *gin.Engine {
	r := gin.New()

	r.Use(middleware.LogrusLogger(logrus.StandardLogger()), gin.Recovery())

	return r
}
