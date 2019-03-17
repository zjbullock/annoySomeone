package service

type Mean interface {
	SendMean() (resp *string, err error)
}

type mean struct {
}

func NewMean() Mean {
	return &mean{}
}

func (m *mean) SendMean() (resp *string, err error) {
	return
}
