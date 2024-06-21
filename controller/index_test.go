package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleShortenUrlReturns400WhenBindingFails(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/shorten", nil)

	HandleShortenUrl(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandleShortenUrlReturns200AndShortenedUrlWhenBindingSucceeds(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("POST", "/shorten", strings.NewReader(`{"long":"https://example.com"}`))
	c.Request.Header.Set("Content-Type", "application/json")

	HandleShortenUrl(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "shortened_url")
}

func TestHandleGetLongUrlReturns404WhenUrlNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "shortUrl", Value: "nonexistent"}}

	HandleGetLongUrl(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestHandleGetLongUrlReturns400WhenUrlNotFound(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "shortUrl", Value: "existent"}}

	HandleGetLongUrl(c)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "short URL not found")
}

func TestHealthCheckReturns200AndStatusUp(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/health", nil)

	HealthCheck(c)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `{"status":"UP"}`)
}
