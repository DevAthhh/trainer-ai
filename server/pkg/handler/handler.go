package handler

import (
	"github.com/DevAthhh/trainer-ai/server/pkg/controllers"
	"github.com/gin-gonic/gin"
)

func Handler() *gin.Engine {
	router := gin.Default()

	// Routes
	api := router.Group("api/")
	{
		api.POST("v1/request", controllers.Request)
		api.POST("v1/check", controllers.Check)
	}

	return router
}
