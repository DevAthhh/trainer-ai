package handler

import (
	"time"

	"github.com/DevAthhh/trainer-ai/client/pkg/controllers"
	"github.com/DevAthhh/trainer-ai/client/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Handle() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.LoadHTMLGlob("client/templates/*")
	router.Static("static/", "client/static/")

	router.GET("/", controllers.Index)
	router.GET("/dashboard", middleware.RequireAuth, controllers.Dashboard)
	router.POST("/questions", middleware.RequireAuth, controllers.Questions)
	router.POST("/result", middleware.RequireAuth)

	users := router.Group("/")
	{
		users.GET("/register", controllers.Register)
		users.POST("/register", controllers.RegisterPost)

		users.GET("/login", controllers.Login)
		users.POST("/login", controllers.LoginPost)

		users.GET("/logout", controllers.Logout)
	}

	return router
}
