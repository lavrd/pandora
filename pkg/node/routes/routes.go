package routes

import (
	"github.com/spacelavr/pandora/pkg/node/routes/handlers"
	"github.com/spacelavr/pandora/pkg/utils/http"
)

var (
	// Routes
	Routes = []http.Route{
		{Path: "/health", Method: http.GET, Middleware: nil, Handler: handlers.HealthH},
		{Path: "/member/create", Method: http.POST, Middleware: nil, Handler: handlers.MemberCreateH},
		{Path: "/member/fetch", Method: http.POST, Middleware: nil, Handler: handlers.MemberFetchH},
		{Path: "/cert/issue", Method: http.POST, Middleware: nil, Handler: handlers.CertIssueH},
		{Path: "/cert/view", Method: http.POST, Middleware: nil, Handler: handlers.CertViewH},
		{Path: "/cert/verify", Method: http.POST, Middleware: nil, Handler: handlers.CertVerifyH},
		{Path: "/blockchain", Method: http.GET, Middleware: nil, Handler: handlers.BlockchainH},
		{Path: "/dashboard", Method: http.GET, Middleware: nil, Handler: handlers.DashboardH},
	}
)
