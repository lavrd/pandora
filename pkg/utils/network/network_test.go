package network_test

import (
	"testing"

	"github.com/spacelavr/pandora/pkg/utils/network"
	"github.com/stretchr/testify/assert"
)

func TestPortWithSemicolon(t *testing.T) {
	cases := []struct {
		Endpoint string
		Port     string
	}{
		{
			Endpoint: "localhost:8080",
			Port:     ":8080",
		},
		{
			Endpoint: "localhost:7777",
			Port:     ":7777",
		},
	}

	for _, c := range cases {
		t.Run(c.Endpoint, func(t *testing.T) {
			assert.Equal(t, c.Port, network.PortWithSemicolon(c.Endpoint))
		})
	}
}
