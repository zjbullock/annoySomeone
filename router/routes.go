package router

import (
	"annoySomeone/router/handlers"
	"context"
	"net/http"
)

var path = "/mean"

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func Routes(ctx context.Context) []route {
	return []route{
		{
			Name:        "Mean",
			Method:      http.MethodPost,
			Pattern:     path,
			HandlerFunc: handlers.BeMean(ctx),
		},
	}
}
