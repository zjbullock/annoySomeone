package router

import (
	"context"
	"github.com/gorilla/mux"
)

func NewRouter(ctx context.Context) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes(ctx) {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
