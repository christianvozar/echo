# Echo

Implementation of the Echo protocol as outlined by [RFC 862](https://tools.ietf.org/html/rfc862). This is implemented as a package in Go and is relatively orthogonal to the net/http package.

Much of the code here is taken from the example in the net package. The purpose here is to illustrate how one might implement this RFC as a package similar to the net/http package for a long-dead protocol standard.

# RFC 862

## Echo Protocol

This RFC specifies a standard for the ARPA Internet community. Hosts on the ARPA Internet that choose to implement an Echo Protocol are expected
to adopt and implement this standard.

A very useful debugging and measurement tool is an echo service. An echo service simply sends back to the originating source any data it receives.

### TCP Based Echo Service

One echo service is defined as a connection based application on TCP. A server listens for TCP connections on TCP port 7. Once a connection is established any data received is sent back. This continues until the calling user terminates the connection.

### UDP Based Echo Service

Another echo service is defined as a datagram based application on UDP. A server listens for UDP datagrams on UDP port 7. When a datagram is received, the data from it is sent back in an answering datagram.

# License and Authors

[Christian Vozar](https://twitter.com/christianvozar) <christian@rogueethic.com>

Copyright © 2015-2016 Christian R. Vozar of [Rogue Ethic](https://github.com/rogueethic)

Use of this source code is governed by a MIT license that can be found in the [LICENSE](LICENSE.markdown) file.
