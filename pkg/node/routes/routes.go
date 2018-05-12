package routes

import (
	"github.com/spacelavr/pandora/pkg/node/routes/handlers"
	"github.com/spacelavr/pandora/pkg/utils/http"
)

var (
	// Routes
	Routes = []http.Route{
		{Path: "/health", Method: http.GET, Middleware: nil, Handler: handlers.HealthH},
		{Path: "/account/create", Method: http.POST, Middleware: nil, Handler: handlers.CandidateH},
		{Path: "/account/fetch", Method: http.POST, Middleware: nil, Handler: handlers.AccountFetchH},
		{Path: "/cert/issue", Method: http.POST, Middleware: nil, Handler: handlers.CertificateIssueH},
		{Path: "/cert/view", Method: http.POST, Middleware: nil, Handler: handlers.CertificateViewH},
		{Path: "/cert/verify", Method: http.POST, Middleware: nil, Handler: handlers.CertificateVerifyH},
		{Path: "/blockchain", Method: http.GET, Middleware: nil, Handler: handlers.BlockchainH},
		{Path: "/dashboard", Method: http.GET, Middleware: nil, Handler: handlers.DashboardH},
	}
)
