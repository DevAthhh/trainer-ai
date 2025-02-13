package main

import (
	"os"

	"github.com/DevAthhh/trainer-ai/server/pkg/handler"
	"github.com/DevAthhh/trainer-ai/server/pkg/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := handler.Handler()
	router.Run(":" + os.Getenv("PORT"))
}
