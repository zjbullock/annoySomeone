package repository

import (
	"annoySomeone/model"
	"encoding/json"
	"fmt"
	"github.com/juju/loggo"
	"github.com/pkg/errors"
	"net/http"
)

type EffOff interface {
	GetMeanStatement(who model.Who) (*string, error)
}

type effOff struct {
	log    loggo.Logger
	client http.Client
	url    string
}

var fmtRequest = http.NewRequest

func NewEffOff(l loggo.Logger, client http.Client, url string) EffOff {
	return &effOff{
		log:    l,
		client: client,
		url:    url,
	}
}

func (e *effOff) GetMeanStatement(who model.Who) (*string, error) {
	e.log.Infof("Repository - EffOff - Formatting Get Request")
	req, err := fmtRequest(http.MethodGet, formatMeanRequest(who, e.url), nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error formatting request")
	}

	req.Header.Set("Accept", "application/json")
	e.log.Infof("Repository - EffOff - Sending Get Request")
	resp, err := e.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "error contacting foaas")
	}
	if resp.StatusCode != 200 {
		mean := "I couldn't get a mean text cuz foaas is being dumb.  My peepee hurt ;("
		return &mean, nil
	}
	e.log.Infof("Repository - EffOff - Successfully Get a Response")
	defer resp.Body.Close()
	var m model.EffOff

	e.log.Infof("Repository - EffOff - Now Decoding Response Body")
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding response into EffOff struct")
	}
	mean := fmt.Sprint(m.Message, m.Subtitle)
	e.log.Infof("Repository - EffOff - Got Response %s", mean)
	return &mean, nil
}

func formatMeanRequest(who model.Who, url string) string {
	return fmt.Sprintf("%s/%s/%s/%s", url, *who.Point, who.To, who.From)
}
