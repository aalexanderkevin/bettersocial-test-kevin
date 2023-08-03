package controller

import (
	"github.com/gin-gonic/gin"
)

func (h *httpServer) setupRouting() {
	router := h.engine

	router.GET("/ping", func(context *gin.Context) {
		context.String(200, "Ok")
	})

	// API
	V1 := router.Group(h.config.Service.Path.V1)
	V1.GET("/user/:username", h.controllers.user.CheckUsername)

	// router.POST("/todos", h.controllers.todo.CreateTodo)
	// router.GET("/todos/:id", h.controllers.todo.GetTodo)
	// router.PUT("/todos/:id", h.controllers.todo.UpdateTodo)
	// router.DELETE("/todos/:id", h.controllers.todo.DeleteTodo)
}
