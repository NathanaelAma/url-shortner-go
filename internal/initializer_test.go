package internal

import (
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestInitializeServer(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(currentFile))
	r := InitializeServer(baseDir)

	req, _ := http.NewRequest("GET", "/", http.NoBody)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "URL Shortener")

	// Add more tests for other routes
	req, _ = http.NewRequest("GET", "/health", http.NoBody)
	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Contains(t, resp.Body.String(), "UP")
}
