package handler

import (
	"github.com/DevAthhh/trainer-ai/server/pkg/controllers"
	"github.com/DevAthhh/trainer-ai/server/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func Handler() *gin.Engine {
	router := gin.Default()

	// Routes
	router.POST("/request", middleware.RequireAuth, controllers.Request) // Getting issues
	router.POST("/check", middleware.RequireAuth, controllers.Check)     // Results of test

	users := router.Group("/")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
		users.POST("/logout", controllers.Logout)
	}

	return router
}
