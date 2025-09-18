package services

import "my-go-practice/pkg/models"

type CustomerServiceImpl struct {
	customers map[int]*models.Customer
}

func NewCustomerServiceImpl() *CustomerServiceImpl {
	return &CustomerServiceImpl{
		customers: map[int]*models.Customer{
			1: {Id: 1, Name: "Alice", Balance: 1000.0, Age: 30},
			2: {Id: 2, Name: "Bob", Balance: 500.0, Age: 22},
			3: {Id: 3, Name: "Charlie", Balance: 200.0, Age: 17},
		},
	}
}
