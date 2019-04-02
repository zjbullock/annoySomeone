package service

import (
	"annoySomeone/model"
	"annoySomeone/repository"
	"annoySomeone/repository/helper"
	"github.com/juju/loggo"
	"github.com/pkg/errors"
)

type Milk interface {
	SendMilk(who model.Who) (resp *string, err error)
}

type milk struct {
	log   loggo.Logger
	wally repository.Wally
	sms   repository.SMS
}

func NewMilk(l loggo.Logger, wally repository.Wally, sms repository.SMS) Milk {
	return &milk{
		log:   l,
		wally: wally,
		sms:   sms,
	}
}

const greatValue = "Great-Value-Whole-Milk-1-Gallon-128-Fl-Oz/10450114"

func (m *milk) SendMilk(who model.Who) (resp *string, err error) {
	secrets, err := helper.GetSecrets(m.log)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting secrets")
	}

	milk, err := m.wally.GetMilkPrice(greatValue, secrets.Wally, who.From)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting milk price")
	}

	resp, err = m.sms.SendText(who.Number, *milk, secrets)
	if err != nil {
		return nil, errors.Wrapf(err, "error sending milk text")
	}
	return
}
