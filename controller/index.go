package controller

import (
	"github.com/NathanaelAma/url-shortener/service"
	"net/http"

	"github.com/NathanaelAma/url-shortener/model"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func ShortenUrl(c *gin.Context) {
	c.HTML(http.StatusOK, "page2.html", nil)
}

func HandleShortenUrl(c *gin.Context) {
	var url model.Url

	if err := c.BindJSON(&url); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shortenedUrl, err := service.ShortenUrl(url)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"shortened_url": shortenedUrl.Short})
}

func HandleGetLongUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	longUrl, err := service.GetLongUrl(shortUrl)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"long_url": longUrl})
}
