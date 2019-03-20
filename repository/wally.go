package repository

import (
	"encoding/json"
	"fmt"
	"github.com/juju/loggo"
	"github.com/pkg/errors"
	"net/http"
)

type Wally interface {
	GetMilkPrice(item, wallyKey string) (*string, error)
}

type wally struct {
	log    loggo.Logger
	client http.Client
	url    string
}

func NewWally(l loggo.Logger, client http.Client, url string) Wally {
	return &wally{
		log:    l,
		client: client,
		url:    url,
	}
}

func (w *wally) GetMilkPrice(item, wallyKey string) (*string, error) {
	w.log.Infof("Repository - Wally - Formatting Get Request")
	req, err := fmtRequest(http.MethodGet, formatMilkRequest(w.url, item, wallyKey), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error formatting request")
	}
	w.log.Infof("Repository - EffOff - Sending Get Request")
	resp, err := w.client.Do(req)
	if resp.StatusCode != 200 {
		milk := "I couldn't get the price of goat milk.  My peepee hurt ;("
		return &milk, nil
	}
	w.log.Infof("Reposiotry - Wally - Successfully Get a Response")
	defer resp.Body.Close()
	var m map[string]interface{}

	w.log.Infof("Repository - Wally - Now Decoding Response Body")
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding response into map")
	}

	milk := string(fmt.Sprintf(`The price of "%s" is $%.2f`, m["name"], m["salePrice"]))
	w.log.Infof("Repository - Wally - Got Response %s", milk)
	return &milk, nil

}

func formatMilkRequest(url, item, wally string) string {
	return fmt.Sprintf("%s/%s?apiKey=%s", url, item, wally)
}
