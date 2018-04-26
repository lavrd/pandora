package routes

import (
	"github.com/spacelavr/pandora/pkg/api/routes/handlers"
	"github.com/spacelavr/pandora/pkg/utils/http"
	"github.com/spacelavr/pandora/pkg/utils/http/middleware"
)

var (
	// Routes
	Routes = []http.Route{
		{Path: "/health", Method: http.GET, Middleware: nil, Handler: handlers.HealthH},

		{Path: "/signin", Method: http.POST, Middleware: nil, Handler: handlers.SignInH},
		{Path: "/signup", Method: http.POST, Middleware: nil, Handler: handlers.SignUpH},
		{Path: "/recovery", Method: http.POST, Middleware: nil, Handler: handlers.AccountRecoveryH},
		{Path: "/account", Method: http.GET, Middleware: []http.Middleware{middleware.Authenticate}, Handler: handlers.AccountFetchH},

		{Path: "/issue", Method: http.POST, Middleware: []http.Middleware{middleware.Issuer, middleware.Authenticate}, Handler: handlers.CertificateIssueH},
		{Path: "/view", Method: http.POST, Middleware: []http.Middleware{middleware.Authenticate}, Handler: handlers.CertificateViewH},
		{Path: "/verify", Method: http.POST, Middleware: []http.Middleware{middleware.Authenticate}, Handler: handlers.CertificateVerifyH},
	}
)
