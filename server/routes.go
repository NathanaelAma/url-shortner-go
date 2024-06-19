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

	// Ping test
	r.GET("/", controller.Index)

	return r
}
