package services

import "my-customer-app/pkg/models"

// CustomerService defines operations on customers
type CustomerService interface {
	GetCustomer(id int) *models.Customer
	UpdateBalance(id int, amount float64)
	PrintCustomer(id int)
}
