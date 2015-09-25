// Copyright © 2015 Rogue Ethic
// Use of this source code is governed by a BSD-style license that can be found
// in the LICENSE.markdown file.

// Much of the code here is taken from the example in the net package. The
// purpose here is to illustrate how one might implement this RFC as a package
// similar to the net/http package.

// Echo server. See RFC 862
// https://tools.ietf.org/html/rfc862

package echo

import (
	"io"
	"net"
)

// Server defines parameters for running a Echo server.
type Server struct {
	Addr string // TCP address to listen on, RFC 862 default of ":7" if empty
}

// ListenAndServe listens on the TCP network address srv.Addr and handles
// requests on incoming connections.
func (srv *Server) ListenAndServe() error {
	addr := srv.Addr
	if addr == "" {
		addr = ":13"
	}
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer ln.Close()

	for {
		// Wait for connection.
		conn, err := ln.Accept()
		if err != nil {
			return err
		}
		// Handle connection in new goroutine.
		// Loop returns to accepting so multiple connections are served concurrently.
		go func(c net.Conn) {
			// Echo all incoming data.
			io.Copy(c, c)
			// Shut down the connection.
			c.Close()
		}(conn)
	}
}

// ListenAndServe listens on the TCP network address addr and serves Echo
// protocol requests in a new goroutine.
func ListenAndServe(addr string) error {
	server := &Server{Addr: addr}
	return server.ListenAndServe()
}
