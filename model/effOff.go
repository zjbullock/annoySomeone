package model

import (
	"encoding/json"
	"fmt"
)

type EffOff struct {
	Message  string `json:"message"`
	Subtitle string `json:"subtitle"`
}

func (e *EffOff) String() string {
	jsn, _ := json.Marshal(&e)
	return fmt.Sprintf(string(jsn))
}
