package main

import (
	"annoySomeone/global"
	"annoySomeone/router"
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/juju/loggo"
	"net/http"
)

var (
	l   loggo.Logger
	ctx = context.Background()
)

func init() {

}

func main() {
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"POST"})
	router := router.NewRouter(ctx)
	fmt.Printf("Listening on port %s\n", global.PORT)
	l.Criticalf(http.ListenAndServe(global.PORT, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)).Error())
}
