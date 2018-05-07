package routes

import (
	"github.com/spacelavr/pandora/pkg/node/routes/handlers"
	"github.com/spacelavr/pandora/pkg/utils/http"
)

var (
	// Routes
	Routes = []http.Route{
		{Path: "/health", Method: http.GET, Middleware: nil, Handler: handlers.HealthH},
		{Path: "/account/new", Method: http.POST, Middleware: nil, Handler: handlers.NewAccountH},
		{Path: "/account/fetch", Method: http.GET, Middleware: nil, Handler: handlers.FetchAccountH},
		{Path: "/cert/new", Method: http.POST, Middleware: nil, Handler: handlers.NewCertificateH},
		{Path: "/cert/fetch", Method: http.GET, Middleware: nil, Handler: handlers.FetchCertificateH},
		{Path: "/cert/verify", Method: http.GET, Middleware: nil, Handler: handlers.VerifyCertificateH},
	}
)
