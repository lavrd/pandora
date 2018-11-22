package routes

import (
	"pandora/pkg/node/routes/handlers"
	"pandora/pkg/utils/http"
)

var (
	// SubRoutes
	SubRoutes = http.SubRoutes{{
		Prefix:     "/member",
		Routes:     handlers.MemberRoutes,
		Middleware: nil,
	}, {
		Prefix:     "/cert",
		Routes:     handlers.CertRoutes,
		Middleware: nil,
	}, {
		Prefix:     "/dashboard",
		Routes:     handlers.DashboardRoutes,
		Middleware: nil,
	}, {
		Prefix:     "/blockchain",
		Routes:     handlers.BlockchainRoutes,
		Middleware: nil,
	}}
)
