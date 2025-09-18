package services

import (
	"fmt"
	"my-customer-app/pkg/models"
	"my-customer-app/pkg/utils"
)

// CustomerServiceImpl is the implementation of CustomerService
type CustomerServiceImpl struct {
	// its like in java Map<int, Customer> customers;
	customers map[int]*models.Customer
}

/*
📌 What *models.Customer means

It means “a pointer to a Customer”, i.e., it stores the memory address where the Customer is kept.

✅ Two ways this can work in Go
➤ 1st way — Explicitly using &models.Customer{...}
customers := map[int]*models.Customer{
    1: &models.Customer{Id: 1, Name: "Alice", Balance: 1000.0, Age: 30},
    2: &models.Customer{Id: 2, Name: "Bob", Balance: 500.0, Age: 25},
    3: &models.Customer{Id: 3, Name: "Charlie", Balance: 200.0, Age: 17},
}


✔ Here:

You explicitly create a Customer object and take its address using &.

*models.Customer expects an address → you give it &models.Customer{...}.

➤ 2nd way — Go's shorthand initialization
customers := map[int]*models.Customer{
    1: {Id: 1, Name: "Alice", Balance: 1000.0, Age: 30},
    2: {Id: 2, Name: "Bob", Balance: 500.0, Age: 25},
    3: {Id: 3, Name: "Charlie", Balance: 200.0, Age: 17},
}


✔ Here:

Go sees that the map expects *models.Customer, so it automatically takes the address of each initialized struct.

This is a convenient shortcut Go provides—you don’t need to explicitly write &.

✅ So what’s happening in both?
Way	Description
1st way	You explicitly create a Customer struct and pass its address using &.
2nd way	Go automatically creates the struct and passes its address, because the map’s value type is *models.Customer.
📌 Important Takeaways

*models.Customer → means the map stores addresses of Customer objects.

You can give it an address using &models.Customer{...}.

Or Go can automatically take the address if you just write {Id:..., Name:...}.

This is why both versions are valid!

✅ Real-time analogy:

*models.Customer asks: “Give me the location where the Customer lives.”

&models.Customer{...} says: “Here’s the location of this Customer.”

Go’s shorthand says: “I know what you want, so I’ll figure out the location for you!”

import java.util.HashMap;
import java.util.Map;

class Customer {
    int id;
    String name;
    double balance;
    int age;

    public Customer(int id, String name, double balance, int age) {
        this.id = id;
        this.name = name;
        this.balance = balance;
        this.age = age;
    }
}

class CustomerServiceImpl {
    Map<Integer, Customer> customers;

    public CustomerServiceImpl() {
        customers = new HashMap<>();
        customers.put(1, new Customer(1, "Alice", 1000.0, 30));
        customers.put(2, new Customer(2, "Bob", 500.0, 25));
        customers.put(3, new Customer(3, "Charlie", 200.0, 17));
    }
}
*/

// NewCustomerService creates and initializes CustomerServiceImpl
func NewCustomerService() *CustomerServiceImpl {
	utils.Log("Initializing Customer Service")
	return &CustomerServiceImpl{
		customers: map[int]*models.Customer{
			1: {Id: 1, Name: "Alice", Balance: 1000.0, Age: 30},
			2: {Id: 2, Name: "Bob", Balance: 500.0, Age: 25},
			3: {Id: 3, Name: "Charlie", Balance: 200.0, Age: 17},
		},
	}
}

// GetCustomer returns the customer by id
func (s *CustomerServiceImpl) GetCustomer(id int) *models.Customer {
	utils.Log(fmt.Sprintf("Getting customer with ID: %d", id))
	return s.customers[id]
}

// UpdateBalance updates the balance for a customer
func (s *CustomerServiceImpl) UpdateBalance(id int, amount float64) {
	utils.Log(fmt.Sprintf("Updating balance for ID: %d by %.2f", id, amount))
	customer := s.GetCustomer(id)
	if customer != nil {
		customer.Balance += amount
	}
}

// PrintCustomer prints customer details
func (s *CustomerServiceImpl) PrintCustomer(id int) {
	utils.Log(fmt.Sprintf("Printing customer details for ID: %d", id))
	customer := s.GetCustomer(id)
	if customer != nil {
		fmt.Printf("Customer ID: %d, Name: %s, Balance: %.2f, Age: %d\n",
			customer.Id, customer.Name, customer.Balance, customer.Age)
	} else {
		fmt.Println("Customer not found!")
	}
}
