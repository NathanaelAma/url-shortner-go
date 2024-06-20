package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/NathanaelAma/url-shortener/server"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	err := os.Setenv("PORT", "8080")
	if err != nil {
		return
	}
	defer func() {
		err := os.Unsetenv("PORT")
		if err != nil {

		}
	}()

	router := server.SetupRouter("./")
	router.Use(sentrygin.New(sentrygin.Options{}))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "UP", response["status"])
}
