package internal

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"path/filepath"
	"runtime"

	"testing"
	"time"
)

func TestStartServer2(t *testing.T) {
	port := "8081"
	env := "test"

	_, currentFile, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filepath.Dir(currentFile))

	// Set environment variable for the test
	t.Setenv("PORT", port)
	t.Setenv("ENV", env)

	// Run the server in a goroutine
	go func() {
		StartServer(baseDir, port, env)
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
