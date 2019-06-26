package handlers

import (
	"annoySomeone/global"
	"annoySomeone/model"
	"annoySomeone/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/juju/loggo"
	"net/http"
)

var (
	l loggo.Logger
)

func BeMean(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := json.NewDecoder(r.Body)
		var who model.Who
		err := req.Decode(&who)
		if err != nil {
			l.Errorf("Error decoding request body: %v", who)
			return
		}
		resp, err := ctx.Value(global.MeanService).(service.Mean).SendMean(who)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, *resp)
	}
}

func GotMilk(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		if params["zipCode"] == "" {
			l.Errorf("Error getting zipCode")
			return
		}
		zipCode := params["zipCode"]
		l.Infof("zipCode: %v", zipCode)
		resp, err := ctx.Value(global.MilkService).(service.Milk).GetMilk(zipCode)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, *resp)
		return
	}
}

func TextMilk(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := json.NewDecoder(r.Body)
		var who model.Who
		err := req.Decode(&who)
		if err != nil {
			l.Errorf("Error decoding request body: %v", who)
			return
		}
		if len(who.Number) != 10 {
			http.Error(w, fmt.Sprint("Length of phone number is not 10 characters."), http.StatusBadRequest)
			return
		}
		if len(who.ZipCode) == 0 {
			http.Error(w, fmt.Sprint("No Zipcode provided"), http.StatusBadRequest)
			return
		}

		resp, err := ctx.Value(global.MilkService).(service.Milk).SendMilk(who)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, *resp)
		return
	}
}
