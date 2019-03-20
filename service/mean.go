package service

import (
	"annoySomeone/model"
	"annoySomeone/repository"
	"github.com/juju/loggo"
	"github.com/pkg/errors"
	"math/rand"
)

type Mean interface {
	SendMean(who model.Who) (resp *string, err error)
}

type mean struct {
	log loggo.Logger
	eff repository.EffOff
	sms repository.SMS
}

func NewMean(l loggo.Logger, eff repository.EffOff, sms repository.SMS) Mean {
	return &mean{
		log: l,
		eff: eff,
		sms: sms,
	}
}

func (m *mean) SendMean(who model.Who) (resp *string, err error) {
	if who.Point == nil {
		point := getPoint()
		m.log.Infof("Randomly generated point with chosen point: %s", point)
		who.Point = &point
	}
	mean, err := m.eff.GetMeanStatement(who)
	if err != nil {
		return nil, errors.Wrapf(err, "error getting mean statement")
	}

	resp, err = m.sms.SendText(who.Number, *mean, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "error sending mean text")
	}
	return
}

func getPoint() string {
	key := rand.Intn(10)
	points := map[int]string{
		0:  "bus",
		1:  "blackadder",
		2:  "bday",
		3:  "back",
		4:  "equity",
		5:  "think",
		6:  "thinking",
		7:  "xmas",
		8:  "madison",
		9:  "look",
		10: "fts",
	}
	return points[key]
}
