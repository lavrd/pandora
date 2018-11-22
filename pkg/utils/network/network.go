package network

import (
	"fmt"
	"net"
)

const (
	TCP = "tcp"
)

// PortWithSemicolon returns port with semicolon from endpoint
func PortWithSemicolon(endpoint string) string {
	_, port, _ := net.SplitHostPort(endpoint)
	return fmt.Sprintf(":%s", port)
}
