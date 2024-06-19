package server

import (
	"github.com/NathanaelAma/url-shortener/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("./public/*.html")

	// Ping test
	r.GET("/", controller.Index)

	return r
}
