// Copyright © 2015-2016 Christian R. Vozar
// MIT Licensed.

package echo

import (
	"testing"
)

// TestEchoServer tests a simple Echo protocol server.
func TestEchoServer(t *testing.T) {
	err := ListenAndServe(":7")
	if err != nil {
		t.Fatal(err)
	}
}
