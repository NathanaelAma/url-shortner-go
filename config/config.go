package config

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"runtime"
)

func SetupConfig() {

	// Set Gin mode based on the environment variable
	env := os.Getenv("ENV")
	fmt.Println("Running on", env)
	switch env {
	case "development":
		gin.SetMode(gin.DebugMode)
	case "testing":
		gin.SetMode(gin.TestMode)
	case "production":
		gin.SetMode(gin.ReleaseMode)
	default:
		gin.SetMode(gin.DebugMode)
	}

	log.Printf("Gin mode set to %s", gin.Mode())
}

func SetupSentry() {
	dsn := os.Getenv("SENTRY_DSN")
	env := os.Getenv("ENV")

	if err := sentry.Init(sentry.ClientOptions{
		// get from viper otherwise fallback to os
		Dsn:           dsn,
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		Environment:      env,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}

func Runtime() {
	nuCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nuCPU)
	fmt.Printf("Running with %d CPUs\n", nuCPU)
}
