package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/NathanaelAma/url-shortener/routes"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Setenv("PORT", "8080")

	router := routes.SetupRouter("./")
	router.Use(sentrygin.New(sentrygin.Options{}))

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", http.NoBody)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "application/json; charset=utf-8", w.Header().Get("Content-Type"))

	var response map[string]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)
	assert.Equal(t, "UP", response["status"])
}
