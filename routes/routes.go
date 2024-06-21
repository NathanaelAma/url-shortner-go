package routes

import (
	"path/filepath"

	"github.com/NathanaelAma/url-shortener/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(baseDir string) *gin.Engine {
	r := gin.Default()

	templatesPath := filepath.Join(baseDir, "public", "*.html")
	staticPath := filepath.Join(baseDir, "public", "static")
	cssPath := filepath.Join(baseDir, "public", "css")

	r.LoadHTMLGlob(templatesPath)
	r.Static("/static", staticPath)
	r.Static("/css", cssPath)

	r.GET("/", controller.Index)
	r.PUT("/shorten", controller.HandleShortenUrl)
	r.GET("/s/:shortUrl", controller.HandleGetLongUrl)

	r.GET("/health", controller.HealthCheck)

	return r
}
