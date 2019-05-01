package model

type Location struct {
	PostalCode    string `json:"postalCode"`
	ResponseGroup string `json:"responseGroup"`
	StoreMeta     bool   `json:"storeMeta"`
	Plus          bool   `json:"plus"`
}
