package config

import (
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

func SetupConfig() {
	viper.SetConfigFile(".env")

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Set Gin mode based on the environment variable
	env := viper.GetString("ENV")
	osEnv := os.Getenv("ENV")
	fmt.Println("Running on", env)
	fmt.Println("Running on", osEnv)
	switch osEnv {
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
	dsn := viper.GetString("SENTRY_DSN")
	if dsn == "" {
		fmt.Println("SENTRY_DSN is not set")
		return
	}
	dsn = os.Getenv("SENTRY_DSN")
	if err := sentry.Init(sentry.ClientOptions{
		// get from viper otherwise fallback to os
		Dsn:           dsn,
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}
}
