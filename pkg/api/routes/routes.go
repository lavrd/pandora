package routes

import (
	"github.com/spacelavr/pandora/pkg/utils/http"
)

var (
	// Routes
	Routes = []http.Route{
		{Path: "/health", Method: http.MethodGet, Middleware: nil, Handler: HealthH},
	}
)
