package network

import (
	"fmt"
	"net"
)

const (
	TCP = "tcp"
)

func PortWithSemicolon(endpoint string) string {
	_, port, _ := net.SplitHostPort(endpoint)
	return fmt.Sprintf(":%s", port)
}
