package global

type key int

const (
	MeanService key = iota
	MilkService
	PORT string = ":8080"
)
