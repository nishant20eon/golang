# Main All Mix

## 🔹 15 Real-Time Entity Questions

### 1. **Employee Management**

- **Struct**: Employee (`ID`, `Name`, `Salary`)
- **Map**: Map of `ID -> Employee`
- **Constructor**: `NewEmployee`
- **Interface**: `Promotable` with method `Promote() float64`
- **Function**: Add employee to map, get salary, update salary (pointer receiver)

---

### 2. **Banking System**

- **Struct**: BankAccount (`AccountNumber`, `HolderName`, `Balance`)
- **Map**: Map of `AccountNumber -> BankAccount`
- **Constructor**: `NewBankAccount`
- **Interface**: `Transaction` with methods `Deposit(amount float64)` and `Withdraw(amount float64)`
- **Function**: Transfer money between two accounts (multi-param, multi-return)

---

### 3. **Product Inventory**

- **Struct**: Product (`ID`, `Name`, `Price`, `Quantity`)
- **Map**: `ProductID -> Product`
- **Constructor**: `NewProduct`
- **Interface**: `Sellable` with `Sell(qty int) float64`
- **Function**: Calculate total inventory value

---

### 4. **Order Processing**

- **Struct**: Order (`OrderID`, `ProductIDs []int`, `CustomerID`, `TotalAmount`)
- **Map**: `OrderID -> Order`
- **Constructor**: `NewOrder`
- **Interface**: `Payable` with `Pay() bool`
- **Function**: Calculate order total (multi-return: total amount, number of items)

---

### 5. **Customer Management**

- **Struct**: Customer (`ID`, `Name`, `Email`)
- **Map**: `ID -> Customer`
- **Constructor**: `NewCustomer`
- **Interface**: `Notifier` with `Notify(message string)`
- **Function**: Send notification to all customers in the map (single param, value receiver)

---

### 6. **Vehicle Fleet**

- **Struct**: Vehicle (`ID`, `Type`, `Brand`, `Speed`)
- **Map**: `ID -> Vehicle`
- **Constructor**: `NewVehicle`
- **Interface**: `Drivable` with `Drive() string`
- **Function**: Calculate average fleet speed (multi-return: avg speed, total vehicles)

---

### 7. **Library System**

- **Struct**: Book (`ID`, `Title`, `Author`, `AvailableCopies`)
- **Map**: `ID -> Book`
- **Constructor**: `NewBook`
- **Interface**: `Borrowable` with `Borrow() bool` and `ReturnBook()`
- **Function**: Borrow multiple books by ID (multi-param, multi-return: success count, failed count)

---

### 8. **Movie Rental**

- **Struct**: Movie (`ID`, `Title`, `Genre`, `Available`)
- **Map**: `ID -> Movie`
- **Constructor**: `NewMovie`
- **Interface**: `Rentable` with `Rent() bool`
- **Function**: Rent multiple movies (single param: slice of IDs, multi-return: rented, failed)

---

### 9. **School Management**

- **Struct**: Student (`ID`, `Name`, `Grade`)
- **Map**: `ID -> Student`
- **Constructor**: `NewStudent`
- **Interface**: `Promotable` with `Promote() string`
- **Function**: Promote multiple students based on grade (pointer receiver)

---

### 🔟 **Ticket Booking**

- **Struct**: Ticket (`TicketID`, `Event`, `Price`, `Booked`)
- **Map**: `TicketID -> Ticket`
- **Constructor**: `NewTicket`
- **Interface**: `Bookable` with `Book() bool`
- **Function**: Calculate total revenue of booked tickets (single return)

---

### 1️⃣1️⃣ **E-Commerce Cart**

- **Struct**: CartItem (`ProductID`, `Quantity`, `Price`)
- **Map**: `UserID -> []CartItem`
- **Constructor**: `NewCartItem`
- **Interface**: `Purchasable` with `Checkout() float64`
- **Function**: Calculate total cart value (multi-param: userID, tax rate)

---

### 1️⃣2️⃣ **Hotel Room Booking**

- **Struct**: Room (`RoomNumber`, `Type`, `Price`, `Occupied`)
- **Map**: `RoomNumber -> Room`
- **Constructor**: `NewRoom`
- **Interface**: `Reservable` with `Reserve() bool`
- **Function**: Reserve multiple rooms (slice of numbers, multi-return: reserved, failed)

---

### 1️⃣3️⃣ **Flight Management**

- **Struct**: Flight (`FlightNo`, `Source`, `Destination`, `SeatsAvailable`)
- **Map**: `FlightNo -> Flight`
- **Constructor**: `NewFlight`
- **Interface**: `Schedulable` with `BookSeat() bool`
- **Function**: Book seats for multiple flights (multi-param, multi-return)

---

### 1️⃣4️⃣ **Warehouse Stock**

- **Struct**: Item (`ID`, `Name`, `Stock`)
- **Map**: `ID -> Item`
- **Constructor**: `NewItem`
- **Interface**: `Stockable` with `AddStock(qty int)` and `ReduceStock(qty int) bool`
- **Function**: Update stock for multiple items at once (multi-param)

---

### 1️⃣5️⃣ **Restaurant Menu**

- **Struct**: Dish (`ID`, `Name`, `Price`)
- **Map**: `ID -> Dish`
- **Constructor**: `NewDish`
- **Interface**: `Orderable` with `Order(qty int) float64`
- **Function**: Calculate total price of multiple dishes (map of dishID -> qty, multi-return: totalPrice, failedDishes)

✅ These 15 exercises cover:

- **Structs + Maps** ✅
- **Constructor functions (factory pattern)** ✅
- **Interfaces with value/pointer receiver methods** ✅
- **Functions with single/multi parameters & single/multi return** ✅
- **Different real-world entities** ✅
- **Designing multi-package modular code** ✅

## 📂 Project Structure

```
company/
├── main.go
├── employee/
│   ├── employee.go
│   └── interface.go
├── bank/
│   ├── account.go
│   └── interface.go
├── product/
│   ├── product.go
│   └── interface.go
├── order/
│   ├── order.go
│   └── interface.go
├── customer/
│   ├── customer.go
│   └── interface.go
├── vehicle/
│   ├── vehicle.go
│   └── interface.go
├── book/
│   ├── book.go
│   └── interface.go
└── movie/
    ├── movie.go
    └── interface.go

```

Each folder represents a **real-world entity**, contains:

- **Struct**
- **Constructor (NewEntity)**
- **Interface**
- **Methods (value & pointer receiver)**

`main.go` will **import all packages**, create **maps of entities**, and call functions demonstrating **single/multi-param, single/multi-return**.

---

## 1️⃣ Employee Package Example

### employee/employee.go

```go
package employee

import "fmt"

// Employee struct
type Employee struct {
	ID     int
	Name   string
	Salary float64
}

// Constructor
func NewEmployee(id int, name string, salary float64) *Employee {
	return &Employee{
		ID:     id,
		Name:   name,
		Salary: salary,
	}
}

// Map to store employees
var Employees = make(map[int]*Employee)

// Add employee to map
func AddEmployee(emp *Employee) {
	Employees[emp.ID] = emp
}

// Pointer receiver method
func (e *Employee) GiveRaise(amount float64) {
	e.Salary += amount
}

// Value receiver method
func (e Employee) Display() {
	fmt.Println("Employee:", e.ID, e.Name, e.Salary)
}

```

### employee/interface.go

```go
package employee

// Promotable interface
type Promotable interface {
	GiveRaise(amount float64)
	Display()
}

```

---

## 2️⃣ Bank Package Example

### bank/account.go

```go
package bank

import "fmt"

// BankAccount struct
type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

// Constructor
func NewBankAccount(accNo, holder string, balance float64) *BankAccount {
	return &BankAccount{
		AccountNumber: accNo,
		HolderName:    holder,
		Balance:       balance,
	}
}

// Map of accounts
var Accounts = make(map[string]*BankAccount)

// Add account
func AddAccount(acc *BankAccount) {
	Accounts[acc.AccountNumber] = acc
}

// Pointer receiver method
func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

// Pointer receiver method
func (b *BankAccount) Withdraw(amount float64) bool {
	if b.Balance >= amount {
		b.Balance -= amount
		return true
	}
	return false
}

// Value receiver method
func (b BankAccount) Display() {
	fmt.Println("Account:", b.AccountNumber, b.HolderName, b.Balance)
}

```

### bank/interface.go

```go
package bank

// Transaction interface
type Transaction interface {
	Deposit(amount float64)
	Withdraw(amount float64) bool
	Display()
}

```

---

## 3️⃣ Main.go Example (Demo All)

```go
package main

import (
	"company/employee"
	"company/bank"
	"fmt"
)

func main() {
	// Employee Example
	emp1 := employee.NewEmployee(1, "Nishant", 75000)
	emp2 := employee.NewEmployee(2, "Rahul", 65000)
	employee.AddEmployee(emp1)
	employee.AddEmployee(emp2)

	// Give raise
	emp1.GiveRaise(5000)
	emp2.GiveRaise(3000)

	// Display employees
	for _, e := range employee.Employees {
		e.Display()
	}

	// Bank Example
	acc1 := bank.NewBankAccount("ACC001", "Nishant", 100000)
	acc2 := bank.NewBankAccount("ACC002", "Rahul", 50000)
	bank.AddAccount(acc1)
	bank.AddAccount(acc2)

	acc1.Deposit(20000)
	success := acc2.Withdraw(10000)
	fmt.Println("Withdrawal success for ACC002:", success)

	for _, a := range bank.Accounts {
		a.Display()
	}
}

```

---

✅ Features demonstrated:

- **Struct + constructor** ✅
- **Map to store entities** ✅
- **Interface with pointer & value receiver** ✅
- **Function calls (single param, multi param, single/multi return)** ✅
- **Multi-package structure** ✅

## 4️⃣ Product Package

### product/product.go

```go
package product

import "fmt"

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
}

// Constructor
func NewProduct(id int, name string, price float64, qty int) *Product {
	return &Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: qty,
	}
}

// Map to store products
var Products = make(map[int]*Product)

// Add product to map
func AddProduct(p *Product) {
	Products[p.ID] = p
}

// Pointer receiver: sell product
func (p *Product) Sell(qty int) float64 {
	if qty > p.Quantity {
		qty = p.Quantity
	}
	p.Quantity -= qty
	return float64(qty) * p.Price
}

// Value receiver: display
func (p Product) Display() {
	fmt.Println("Product:", p.ID, p.Name, p.Price, "Qty:", p.Quantity)
}

```

### product/interface.go

```go
package product

type Sellable interface {
	Sell(qty int) float64
	Display()
}

```

---

## 5️⃣ Order Package

### order/order.go

```go
package order

import "fmt"

type Order struct {
	OrderID    int
	ProductIDs []int
	CustomerID int
	TotalAmount float64
}

func NewOrder(orderID int, products []int, customerID int, total float64) *Order {
	return &Order{
		OrderID:    orderID,
		ProductIDs: products,
		CustomerID: customerID,
		TotalAmount: total,
	}
}

var Orders = make(map[int]*Order)

func AddOrder(o *Order) {
	Orders[o.OrderID] = o
}

// Calculate total number of products (multi-return)
func (o *Order) Summary() (int, float64) {
	return len(o.ProductIDs), o.TotalAmount
}

// Display
func (o Order) Display() {
	fmt.Println("Order:", o.OrderID, "Customer:", o.CustomerID, "Total:", o.TotalAmount)
}

```

### order/interface.go

```go
package order

type Payable interface {
	Summary() (int, float64)
	Display()
}

```

---

## 6️⃣ Customer Package

### customer/customer.go

```go
package customer

import "fmt"

type Customer struct {
	ID    int
	Name  string
	Email string
}

func NewCustomer(id int, name, email string) *Customer {
	return &Customer{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

var Customers = make(map[int]*Customer)

func AddCustomer(c *Customer) {
	Customers[c.ID] = c
}

// Value receiver: notify
func (c Customer) Notify(msg string) {
	fmt.Printf("Notification for %s: %s\n", c.Name, msg)
}

// Display
func (c Customer) Display() {
	fmt.Println("Customer:", c.ID, c.Name, c.Email)
}

```

### customer/interface.go

```go
package customer

type Notifier interface {
	Notify(msg string)
	Display()
}

```

---

## 7️⃣ Vehicle Package

### vehicle/vehicle.go

```go
package vehicle

import "fmt"

type Vehicle struct {
	ID    int
	Type  string
	Brand string
	Speed float64
}

func NewVehicle(id int, vtype, brand string, speed float64) *Vehicle {
	return &Vehicle{
		ID:    id,
		Type:  vtype,
		Brand: brand,
		Speed: speed,
	}
}

var Fleet = make(map[int]*Vehicle)

func AddVehicle(v *Vehicle) {
	Fleet[v.ID] = v
}

// Pointer receiver: drive
func (v *Vehicle) Drive() string {
	return fmt.Sprintf("%s %s is driving at %.2f km/h", v.Type, v.Brand, v.Speed)
}

// Value receiver: display
func (v Vehicle) Display() {
	fmt.Println("Vehicle:", v.ID, v.Type, v.Brand, "Speed:", v.Speed)
}

```

### vehicle/interface.go

```go
package vehicle

type Drivable interface {
	Drive() string
	Display()
}

```

---

## 8️⃣ Book Package

### book/book.go

```go
package book

import "fmt"

type Book struct {
	ID             int
	Title          string
	Author         string
	AvailableCopies int
}

func NewBook(id int, title, author string, copies int) *Book {
	return &Book{
		ID:             id,
		Title:          title,
		Author:         author,
		AvailableCopies: copies,
	}
}

var Library = make(map[int]*Book)

func AddBook(b *Book) {
	Library[b.ID] = b
}

// Pointer receiver: borrow book
func (b *Book) Borrow() bool {
	if b.AvailableCopies > 0 {
		b.AvailableCopies--
		return true
	}
	return false
}

// Pointer receiver: return book
func (b *Book) ReturnBook() {
	b.AvailableCopies++
}

// Display
func (b Book) Display() {
	fmt.Println("Book:", b.ID, b.Title, b.Author, "Available:", b.AvailableCopies)
}

```

### book/interface.go

```go
package book

type Borrowable interface {
	Borrow() bool
	ReturnBook()
	Display()
}

```

---

## 9️⃣ Movie Package

### movie/movie.go

```go
package movie

import "fmt"

type Movie struct {
	ID        int
	Title     string
	Genre     string
	Available bool
}

func NewMovie(id int, title, genre string) *Movie {
	return &Movie{
		ID:        id,
		Title:     title,
		Genre:     genre,
		Available: true,
	}
}

var Movies = make(map[int]*Movie)

func AddMovie(m *Movie) {
	Movies[m.ID] = m
}

// Pointer receiver: rent
func (m *Movie) Rent() bool {
	if m.Available {
		m.Available = false
		return true
	}
	return false
}

// Display
func (m Movie) Display() {
	fmt.Println("Movie:", m.ID, m.Title, m.Genre, "Available:", m.Available)
}

```

### movie/interface.go

```go
package movie

type Rentable interface {
	Rent() bool
	Display()
}

```

## 10️⃣ Student Package

### student/student.go

```go
package student

import "fmt"

type Student struct {
	ID    int
	Name  string
	Grade int
}

func NewStudent(id int, name string, grade int) *Student {
	return &Student{
		ID:    id,
		Name:  name,
		Grade: grade,
	}
}

var Students = make(map[int]*Student)

func AddStudent(s *Student) {
	Students[s.ID] = s
}

// Pointer receiver: promote
func (s *Student) Promote() {
	s.Grade++
}

// Value receiver: display
func (s Student) Display() {
	fmt.Println("Student:", s.ID, s.Name, "Grade:", s.Grade)
}

```

### student/interface.go

```go
package student

type Promotable interface {
	Promote()
	Display()
}

```

---

## 11️⃣ Ticket Package

### ticket/ticket.go

```go
package ticket

import "fmt"

type Ticket struct {
	TicketID int
	Event    string
	Price    float64
	Booked   bool
}

func NewTicket(id int, event string, price float64) *Ticket {
	return &Ticket{
		TicketID: id,
		Event:    event,
		Price:    price,
		Booked:   false,
	}
}

var Tickets = make(map[int]*Ticket)

func AddTicket(t *Ticket) {
	Tickets[t.TicketID] = t
}

// Pointer receiver: book ticket
func (t *Ticket) Book() bool {
	if !t.Booked {
		t.Booked = true
		return true
	}
	return false
}

// Display
func (t Ticket) Display() {
	fmt.Println("Ticket:", t.TicketID, t.Event, "Price:", t.Price, "Booked:", t.Booked)
}

```

### ticket/interface.go

```go
package ticket

type Bookable interface {
	Book() bool
	Display()
}

```

---

## 12️⃣ Cart Package

### cart/cart.go

```go
package cart

import "fmt"

type CartItem struct {
	ProductID int
	Quantity  int
	Price     float64
}

func NewCartItem(pid, qty int, price float64) *CartItem {
	return &CartItem{
		ProductID: pid,
		Quantity:  qty,
		Price:     price,
	}
}

var UserCarts = make(map[int][]*CartItem)

// Add item to user cart
func AddToCart(userID int, item *CartItem) {
	UserCarts[userID] = append(UserCarts[userID], item)
}

// Pointer receiver: calculate total price
func (c *CartItem) Total() float64 {
	return float64(c.Quantity) * c.Price
}

// Value receiver: display item
func (c CartItem) Display() {
	fmt.Println("CartItem:", c.ProductID, "Qty:", c.Quantity, "Price:", c.Price)
}

```

### cart/interface.go

```go
package cart

type Purchasable interface {
	Total() float64
	Display()
}

```

---

## 13️⃣ Room Package

### room/room.go

```go
package room

import "fmt"

type Room struct {
	RoomNumber int
	Type       string
	Price      float64
	Occupied   bool
}

func NewRoom(number int, rtype string, price float64) *Room {
	return &Room{
		RoomNumber: number,
		Type:       rtype,
		Price:      price,
		Occupied:   false,
	}
}

var Rooms = make(map[int]*Room)

func AddRoom(r *Room) {
	Rooms[r.RoomNumber] = r
}

// Pointer receiver: reserve room
func (r *Room) Reserve() bool {
	if !r.Occupied {
		r.Occupied = true
		return true
	}
	return false
}

// Value receiver: display
func (r Room) Display() {
	fmt.Println("Room:", r.RoomNumber, r.Type, "Price:", r.Price, "Occupied:", r.Occupied)
}

```

### room/interface.go

```go
package room

type Reservable interface {
	Reserve() bool
	Display()
}

```

---

## 14️⃣ Flight Package

### flight/flight.go

```go
package flight

import "fmt"

type Flight struct {
	FlightNo       string
	Source         string
	Destination    string
	SeatsAvailable int
}

func NewFlight(no, source, dest string, seats int) *Flight {
	return &Flight{
		FlightNo:       no,
		Source:         source,
		Destination:    dest,
		SeatsAvailable: seats,
	}
}

var Flights = make(map[string]*Flight)

func AddFlight(f *Flight) {
	Flights[f.FlightNo] = f
}

// Pointer receiver: book seat
func (f *Flight) BookSeat() bool {
	if f.SeatsAvailable > 0 {
		f.SeatsAvailable--
		return true
	}
	return false
}

// Value receiver: display
func (f Flight) Display() {
	fmt.Println("Flight:", f.FlightNo, f.Source, "->", f.Destination, "Seats:", f.SeatsAvailable)
}

```

### flight/interface.go

```go
package flight

type Schedulable interface {
	BookSeat() bool
	Display()
}

```

---

## 15️⃣ Dish Package

### dish/dish.go

```go
package dish

import "fmt"

type Dish struct {
	ID    int
	Name  string
	Price float64
}

func NewDish(id int, name string, price float64) *Dish {
	return &Dish{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

var Menu = make(map[int]*Dish)

func AddDish(d *Dish) {
	Menu[d.ID] = d
}

// Pointer receiver: order dish
func (d *Dish) Order(qty int) float64 {
	return float64(qty) * d.Price
}

// Value receiver: display
func (d Dish) Display() {
	fmt.Println("Dish:", d.ID, d.Name, d.Price)
}

```

### dish/interface.go

```go
package dish

type Orderable interface {
	Order(qty int) float64
	Display()
}

```

---

✅ Now the **15-entity project is complete** with:

- **Struct + constructor + map**
- **Interface + value/pointer receiver**
- **Functions with single/multi param & return**
- **Multi-package structure**