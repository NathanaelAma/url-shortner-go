package test

import (
	"github.com/NathanaelAma/url-shortener/internal"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"testing"
)

func TestServerSetup(t *testing.T) {
	gin.SetMode(gin.TestMode)

	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(currentFile))

	r := internal.InitializeServer(baseDir)

	req, _ := http.NewRequest("GET", "/", http.NoBody)
	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Fatalf("Expected status code %d but got %d", http.StatusOK, resp.Code)
	}
}
