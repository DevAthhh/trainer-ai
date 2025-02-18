package main

import (
	"os"

	"github.com/DevAthhh/trainer-ai/client/pkg/handler"
	"github.com/DevAthhh/trainer-ai/client/pkg/initializers"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := handler.Handle()
	router.Run(":" + os.Getenv("PORT_CLIENT"))
}
