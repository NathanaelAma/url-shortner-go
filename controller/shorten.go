package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShortenUrl(c *gin.Context) {
	c.HTML(http.StatusOK, "page2.html", nil)
}
