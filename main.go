package main

import (
	"fmt"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/spf13/viper"
	"runtime"
	"url-shortener/config"
)

func main() {

	ConfigRuntime()
	StartServer()
}

func ConfigRuntime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}

func StartServer() {
	fmt.Println("Starting server...")
	config.SetupConfig()
	config.SetupSentry()
	r := setupRouter()
	r.Use(sentrygin.New(sentrygin.Options{}))
	err := r.Run(":" + viper.GetString("PORT"))
	if err != nil {
		return
	}
}
