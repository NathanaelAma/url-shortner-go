package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShortenUrl(c *gin.Context) {
	c.HTML(http.StatusOK, "page2.html", nil)
}
