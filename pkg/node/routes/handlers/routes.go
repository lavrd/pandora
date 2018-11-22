package handlers

import (
	"pandora/pkg/utils/http"
)

var (
	// MemberRoutes
	MemberRoutes = http.Routes{
		{Path: "/create", Method: http.Post, Middleware: nil, Handler: MemberCreateH},
		{Path: "/fetch", Method: http.Post, Middleware: nil, Handler: MemberFetchH},
	}

	// CertRoutes
	CertRoutes = http.Routes{
		{Path: "/issue", Method: http.Post, Middleware: nil, Handler: MemberFetchH},
		{Path: "/view", Method: http.Post, Middleware: nil, Handler: MemberFetchH},
		{Path: "/verify", Method: http.Post, Middleware: nil, Handler: MemberFetchH},
	}

	// DashboardRoutes
	DashboardRoutes = http.Routes{
		{Path: "", Method: http.Get, Middleware: nil, Handler: DashboardH},
	}

	// BlockchainRoutes
	BlockchainRoutes = http.Routes{
		{Path: "", Method: http.Get, Middleware: nil, Handler: BlockchainH},
	}
)
