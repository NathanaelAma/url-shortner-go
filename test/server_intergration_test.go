package test

import (
	"github.com/NathanaelAma/url-shortener/internal"
	"github.com/stretchr/testify/assert"
	"net/http"
	"path/filepath"
	"runtime"

	"testing"
	"time"
)

func TestStartServer2(t *testing.T) {
	port := "8081"

	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(currentFile))

	// Set environment variable for the test
	t.Setenv("PORT", port)

	// Run the server in a goroutine
	go func() {
		internal.StartServer(baseDir, port)
	}()

	// Allow some time for the server to start
	time.Sleep(1 * time.Second)

	// Send a request to the running server
	resp, err := http.Get("http://localhost:" + port + "/")
	if err != nil {
		t.Fatalf("Could not get the response from the server: %v", err)
	}

	// Verify the response
	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
