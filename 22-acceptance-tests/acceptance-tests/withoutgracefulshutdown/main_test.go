package main

import (
	"testing"
	"time"

	acceptancetests "github.com/tiagocorrea/go-graceful-shutdown/acceptance-tests"
	"github.com/tiagocorrea/go-graceful-shutdown/assert"
)

const (
	port = "8081"
	url  = "http://localhost:" + port
)

func TestGracefulShutdown(t *testing.T) {
	cleanup, sendInterrupt, err := acceptancetests.LaunchTestProgram(port)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(cleanup)

	// just check the server works before we shut things down. uncomment to get an error
	// assert.CanGet(t, url)

	// fire off a request , and before it has a chance to respond send SIGTERM
	time.AfterFunc(50*time.Millisecond, func() {
		assert.NoError(t, sendInterrupt())
	})
	// Without graceful shutdown, this would fail. uncomment to get an error
	// assert.CanGet(t, url)

	// after interrupt, the server should be shutdown, and no more requests will work
	assert.CantGet(t, url)
}
