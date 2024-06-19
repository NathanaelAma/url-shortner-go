package main

import (
	"fmt"
	"github.com/NathanaelAma/url-shortener/config"
	"github.com/NathanaelAma/url-shortener/server"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"os"
)

func main() {

	config.Runtime()
	StartServer()
}

func StartServer() {
	fmt.Println("Starting server...")
	config.SetupConfig()
	config.SetupSentry()
	r := server.SetupRouter()
	r.Use(sentrygin.New(sentrygin.Options{}))
	err := r.Run(":" + os.Getenv("PORT"))
	if err != nil {
		return
	}
}
