package handlers

import (
	"annoySomeone/global"
	"annoySomeone/service"
	"context"
	"flag"
	"fmt"
	"net/http"
)

func BeMean(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flag.Parse()

		resp, err := ctx.Value(global.MeanService).(service.Mean).SendMean()
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("content-type", "application/json")
		fmt.Fprintf(w, *resp)
	}
}
