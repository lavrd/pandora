package network

import (
	"net"
)

func IP() string {
	c, _ := net.Dial("udp", "8.8.8.8:80")
	defer c.Close()

	return c.LocalAddr().String()
}
