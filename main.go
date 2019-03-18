package main

import (
	"annoySomeone/global"
	"annoySomeone/repository"
	"annoySomeone/router"
	"annoySomeone/service"
	"context"
	"github.com/gorilla/handlers"
	"github.com/juju/loggo"
	"net/http"
)

var (
	l      loggo.Logger
	ctx    = context.Background()
	client = http.Client{}
)

func init() {
	l.SetLogLevel(loggo.INFO)
	//Create Repositories
	eff := repository.NewEffOff(l, client, "https://www.foaas.com")
	sms := repository.NewSMS(l, client, "https://api.twilio.com")
	//Create Services
	mean := service.NewMean(l, eff, sms)
	ctx = context.WithValue(ctx, global.MeanService, mean)
}

func main() {
	allowedHeaders := handlers.AllowedHeaders([]string{"content-type"})
	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"POST"})
	router := router.NewRouter(ctx)
	l.Infof("Listening on port %s", global.PORT)
	l.Criticalf(http.ListenAndServe(global.PORT, handlers.CORS(allowedOrigins, allowedMethods, allowedHeaders)(router)).Error())
}
