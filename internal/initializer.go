package internal

import (
	"fmt"
	"github.com/NathanaelAma/url-shortener/config"
	"github.com/NathanaelAma/url-shortener/routes"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
)

func StartServer(baseDir, port, env string) {
	fmt.Println("Starting server...")
	config.SetupConfig(env)
	config.SetupSentry()
	r := InitializeServer(baseDir)
	err := r.Run(":" + port)
	if err != nil {
		return
	}
}

func InitializeServer(baseDir string) *gin.Engine {
	r := routes.SetupRouter(baseDir)
	r.Use(sentrygin.New(sentrygin.Options{}))
	return r
}
