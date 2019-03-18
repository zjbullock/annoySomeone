package repository

import (
	"annoySomeone/repository/helper"
	"fmt"
	"github.com/juju/loggo"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type SMS interface {
	SendText(number, message string) (*string, error)
}

type sms struct {
	log    loggo.Logger
	client http.Client
	url    string
}

func NewSMS(l loggo.Logger, client http.Client, url string) SMS {
	return &sms{
		log:    l,
		client: client,
		url:    url,
	}
}

func (s *sms) SendText(number, message string) (*string, error) {
	s.log.Infof("Repository - SMS - Getting Secret Account Info")
	account, err := helper.GetSecrets(s.log)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting secrets")
	}
	msg := url.Values{}
	msg.Set("To", number)
	msg.Set("From", account.Number)
	msg.Set("Body", message)
	msgDataReader := *strings.NewReader(msg.Encode())

	s.log.Infof("Repository - SMS - Formatting Get Request")
	req, err := fmtRequest(http.MethodPost, formatTwilioRequest(s.url, account.SID), &msgDataReader)
	req.SetBasicAuth(account.SID, account.AuthToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	s.log.Infof("Repository - SMS - Making request")
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, errors.Wrapf(err, "error making the request")
	}
	defer resp.Body.Close()
	bod, _ := ioutil.ReadAll(resp.Body)
	smsResponse := string(bod)

	return &smsResponse, nil
}

func formatTwilioRequest(url, sid string) string {
	return fmt.Sprintf("%s/2010-04-01/Accounts/%s/Messages.json", url, sid)
}