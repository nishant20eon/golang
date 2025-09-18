package models

type Order struct {
	Id         int
	CustomerId int
	Amount     float64
	Status     string
}
