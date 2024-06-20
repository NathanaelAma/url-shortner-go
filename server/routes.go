package server

import (
	"github.com/NathanaelAma/url-shortener/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("./public/*.html")
	r.Static("/static", "./public/static")
	r.Static("/css", "./public/css")

	r.GET("/", controller.Index)

	r.GET("/page2", controller.ShortenUrl)
	r.GET("/health", controller.HealthCheck)

	return r
}
