package handlers

import (
	"annoySomeone/global"
	"annoySomeone/model"
	"annoySomeone/service"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

func BeMean(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flag.Parse()

		req := json.NewDecoder(r.Body)
		var who model.Who
		req.Decode(&who)

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
		flag.Parse()

		req := json.NewDecoder(r.Body)
		var who model.Who
		req.Decode(&who)

		resp, err := ctx.Value(global.MilkService).(service.Milk).SendMilk(who)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, *resp)
	}
}
