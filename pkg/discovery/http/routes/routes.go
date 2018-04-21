package routes

import (
	"github.com/spacelavr/pandora/pkg/discovery/http/routes/handlers"
	"github.com/spacelavr/pandora/pkg/utils/http"
)

var (
	// Routes
	Routes = []http.Route{
		{Path: "/blockchain", Method: http.GET, Middleware: nil, Handler: handlers.BlockchainH},
	}
)
