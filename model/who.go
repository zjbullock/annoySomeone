package model

import (
	"encoding/json"
	"fmt"
)

type Who struct {
	Point   *string `json:"point, omitempty"`
	To      string  `json:"to"`
	From    string  `json:"from"`
	Number  string  `json:"phoneNumber"`
	ZipCode string  `json:"zipCode"`
}

func (w *Who) String() string {
	jsn, _ := json.Marshal(&w)
	return fmt.Sprintf(string(jsn))
}
