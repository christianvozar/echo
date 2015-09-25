// Copyright © 2015 Rogue Ethic
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE.markdown file.

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
