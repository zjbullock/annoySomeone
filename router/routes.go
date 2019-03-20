package router

import (
	"annoySomeone/router/handlers"
	"context"
	"net/http"
)

var (
	mean = "/mean"
	milk = "/milk"
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
			Name:        "Mean",
			Method:      http.MethodPost,
			Pattern:     mean,
			HandlerFunc: handlers.BeMean(ctx),
		},
		{
			Name:        "Milk",
			Method:      http.MethodPost,
			Pattern:     milk,
			HandlerFunc: handlers.GotMilk(ctx),
		},
	}
}
