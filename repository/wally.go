package repository

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/juju/loggo"
	"github.com/pkg/errors"
	"net/http"
)

type Wally interface {
	GetMilkPrice(item, wallyKey, from string) (*string, error)
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

const (
	mulk = "Great Value Whole Milk, 1 Gallon, 128 Fl. Oz."
)

func (w *wally) GetMilkPrice(item, wallyKey, from string) (*string, error) {
	w.log.Infof("Repository - Wally - Formatting Get Request")
	req, err := fmtRequest(http.MethodGet, formatMilkRequest(w.url, item), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error formatting request")
	}
	w.log.Infof("Repository - EffOff - Sending Get Request")
	resp, err := w.client.Do(req)
	if resp.StatusCode != 200 {
		milk := "I couldn't get the price of milk.  My peepee hurt ;("
		return &milk, nil
	}
	w.log.Infof("Repository - Wally - Successfully Get a Response")
	defer resp.Body.Close()
	m := make(map[string]interface{})

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	doc.Find(".price-group").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		m["salePrice"] = title
	})
	m["name"] = mulk
	milk := string(fmt.Sprintf(`The price of "%s" is %s, - %s`, m["name"], m["salePrice"], from))
	w.log.Infof("Repository - Wally - Got Response %s", milk)
	return &milk, nil

}

func formatMilkRequest(url, item string) string {
	return fmt.Sprintf("%s/%s", url, item)
}
