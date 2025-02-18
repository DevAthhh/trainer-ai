package main

import (
	"os"

	"github.com/DevAthhh/trainer-ai/server/initializers"
	"github.com/DevAthhh/trainer-ai/server/pkg/handler"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	router := handler.Handler()
	router.Run(":" + os.Getenv("PORT_SERVER"))
}
