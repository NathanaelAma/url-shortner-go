package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/NathanaelAma/url-shortener/server"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestStartServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Setenv("PORT", "8080")
	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(currentFile))

	router := server.SetupRouter(baseDir)
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
