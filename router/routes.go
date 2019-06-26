package router

import (
	"annoySomeone/router/handlers"
	"context"
	"net/http"
)

var (
	mean     = "/mean"
	textMilk = "/milk"
	getMilk  = "/getMilk/{zipCode}"
)

type route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

func Routes(ctx context.Context) []route {
	return []route{
		{
			Name:        "Text Mean",
			Method:      http.MethodPost,
			Pattern:     mean,
			HandlerFunc: handlers.BeMean(ctx),
		},
		{
			Name:        "Text Milk",
			Method:      http.MethodPost,
			Pattern:     textMilk,
			HandlerFunc: handlers.TextMilk(ctx),
		},
		{
			Name:        "Get Milk",
			Method:      http.MethodGet,
			Pattern:     getMilk,
			HandlerFunc: handlers.GotMilk(ctx),
		},
	}
}
