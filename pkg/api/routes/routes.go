package routes

import (
	"github.com/spacelavr/pandora/pkg/utils/http"
	"github.com/spacelavr/pandora/pkg/utils/http/middleware"
)

var (
	// Routes
	Routes = []http.Route{
		{Path: "/health", Method: http.GET, Middleware: nil, Handler: HealthH},

		{Path: "/signin", Method: http.POST, Middleware: nil, Handler: SignInH},
		{Path: "/signup", Method: http.POST, Middleware: nil, Handler: SignUpH},
		{Path: "/recovery", Method: http.POST, Middleware: nil, Handler: AccountRecoveryH},
		{Path: "/account", Method: http.GET, Middleware: []http.Middleware{middleware.Authenticate}, Handler: AccountFetchH},

		{Path: "/create", Method: http.POST, Middleware: []http.Middleware{middleware.Authenticate}, Handler: CertificateCreateH},
		{Path: "/issue", Method: http.GET, Middleware: []http.Middleware{middleware.Authenticate}, Handler: CertificateIssueH},
		{Path: "/verify", Method: http.POST, Middleware: []http.Middleware{middleware.Authenticate}, Handler: CertificateVerifyH},
	}
)
