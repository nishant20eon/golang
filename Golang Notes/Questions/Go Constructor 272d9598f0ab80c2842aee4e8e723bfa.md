# Go Constructor

## 🔹 10 Real-Time Entity Practice Questions on Go Struct + Constructor

### 1. **Employee Struct**

- Create an `Employee` struct with `Name`, `Age`, and `Salary`.
- Write a constructor function `NewEmployee(name string, age int, salary float64) *Employee`.
- Create an employee and print details.

---

### 2. **Bank Account**

- Struct `BankAccount` with `AccountNumber`, `HolderName`, `Balance`.
- Constructor `NewBankAccount(accountNumber, holderName string, balance float64) *BankAccount`.
- Initialize and display account info.

---

### 3. **Car Struct**

- Struct `Car` with `Brand`, `Model`, `Year`.
- Constructor `NewCar(brand, model string, year int) *Car`.
- Create a car and print its info.

---

### 4. **Book Struct**

- Struct `Book` with `Title`, `Author`, `Price`.
- Constructor `NewBook(title, author string, price float64) *Book`.
- Initialize a book and display details.

---

### 5. **Student Struct**

- Struct `Student` with `ID`, `Name`, `Marks`.
- Constructor `NewStudent(id int, name string, marks float64) *Student`.
- Create a student and print marks.

---

### 6. **Product Struct**

- Struct `Product` with `ID`, `Name`, `Price`, `Quantity`.
- Constructor `NewProduct(id int, name string, price float64, quantity int) *Product`.
- Create a product and calculate total value (`Price * Quantity`).

---

### 7. **Order Struct**

- Struct `Order` with `OrderID`, `ProductName`, `Quantity`, `Price`.
- Constructor `NewOrder(orderID int, productName string, quantity int, price float64) *Order`.
- Calculate total order amount.

---

### 8. **Movie Struct**

- Struct `Movie` with `Title`, `Director`, `Rating`.
- Constructor `NewMovie(title, director string, rating float64) *Movie`.
- Print movie details.

---

### 9. **Customer Struct**

- Struct `Customer` with `CustomerID`, `Name`, `Email`.
- Constructor `NewCustomer(id int, name, email string) *Customer`.
- Print customer info.

---

### 🔟 **Laptop Struct**

- Struct `Laptop` with `Brand`, `Model`, `RAM`, `Price`.
- Constructor `NewLaptop(brand, model string, ram int, price float64) *Laptop`.
- Create laptop object and display specs.

---

## 1️⃣ Employee Struct

```go
package main

import "fmt"

type Employee struct {
    Name   string
    Age    int
    Salary float64
}

func NewEmployee(name string, age int, salary float64) *Employee {
    return &Employee{
        Name:   name,
        Age:    age,
        Salary: salary,
    }
}

func main() {
    emp := NewEmployee("Nishant", 28, 75000)
    fmt.Println("Employee:", emp.Name, emp.Age, emp.Salary)
}

```

---

## 2️⃣ Bank Account

```go
package main

import "fmt"

type BankAccount struct {
    AccountNumber string
    HolderName    string
    Balance       float64
}

func NewBankAccount(accountNumber, holderName string, balance float64) *BankAccount {
    return &BankAccount{
        AccountNumber: accountNumber,
        HolderName:    holderName,
        Balance:       balance,
    }
}

func main() {
    account := NewBankAccount("123456789", "Nishant", 100000)
    fmt.Println("Account:", account.AccountNumber, account.HolderName, account.Balance)
}

```

---

## 3️⃣ Car Struct

```go
package main

import "fmt"

type Car struct {
    Brand string
    Model string
    Year  int
}

func NewCar(brand, model string, year int) *Car {
    return &Car{
        Brand: brand,
        Model: model,
        Year:  year,
    }
}

func main() {
    car := NewCar("Toyota", "Corolla", 2020)
    fmt.Println("Car:", car.Brand, car.Model, car.Year)
}

```

---

## 4️⃣ Book Struct

```go
package main

import "fmt"

type Book struct {
    Title  string
    Author string
    Price  float64
}

func NewBook(title, author string, price float64) *Book {
    return &Book{
        Title:  title,
        Author: author,
        Price:  price,
    }
}

func main() {
    book := NewBook("Go Programming", "Nishant", 499)
    fmt.Println("Book:", book.Title, book.Author, book.Price)
}

```

---

## 5️⃣ Student Struct

```go
package main

import "fmt"

type Student struct {
    ID    int
    Name  string
    Marks float64
}

func NewStudent(id int, name string, marks float64) *Student {
    return &Student{
        ID:    id,
        Name:  name,
        Marks: marks,
    }
}

func main() {
    student := NewStudent(101, "Nishant", 88.5)
    fmt.Println("Student:", student.ID, student.Name, student.Marks)
}

```

---

## 6️⃣ Product Struct

```go
package main

import "fmt"

type Product struct {
    ID       int
    Name     string
    Price    float64
    Quantity int
}

func NewProduct(id int, name string, price float64, quantity int) *Product {
    return &Product{
        ID:       id,
        Name:     name,
        Price:    price,
        Quantity: quantity,
    }
}

func (p *Product) TotalValue() float64 {
    return p.Price * float64(p.Quantity)
}

func main() {
    product := NewProduct(1, "Laptop", 50000, 2)
    fmt.Println("Product:", product.Name, "Total Value:", product.TotalValue())
}

```

---

## 7️⃣ Order Struct

```go
package main

import "fmt"

type Order struct {
    OrderID     int
    ProductName string
    Quantity    int
    Price       float64
}

func NewOrder(orderID int, productName string, quantity int, price float64) *Order {
    return &Order{
        OrderID:     orderID,
        ProductName: productName,
        Quantity:    quantity,
        Price:       price,
    }
}

func (o *Order) TotalAmount() float64 {
    return o.Price * float64(o.Quantity)
}

func main() {
    order := NewOrder(101, "Laptop", 3, 45000)
    fmt.Println("Order Total:", order.TotalAmount())
}

```

---

## 8️⃣ Movie Struct

```go
package main

import "fmt"

type Movie struct {
    Title    string
    Director string
    Rating   float64
}

func NewMovie(title, director string, rating float64) *Movie {
    return &Movie{
        Title:    title,
        Director: director,
        Rating:   rating,
    }
}

func main() {
    movie := NewMovie("Inception", "Christopher Nolan", 9.0)
    fmt.Println("Movie:", movie.Title, movie.Director, movie.Rating)
}

```

---

## 9️⃣ Customer Struct

```go
package main

import "fmt"

type Customer struct {
    CustomerID int
    Name       string
    Email      string
}

func NewCustomer(id int, name, email string) *Customer {
    return &Customer{
        CustomerID: id,
        Name:       name,
        Email:      email,
    }
}

func main() {
    customer := NewCustomer(1, "Nishant", "nishant@example.com")
    fmt.Println("Customer:", customer.CustomerID, customer.Name, customer.Email)
}

```

---

## 🔟 Laptop Struct

```go
package main

import "fmt"

type Laptop struct {
    Brand string
    Model string
    RAM   int
    Price float64
}

func NewLaptop(brand, model string, ram int, price float64) *Laptop {
    return &Laptop{
        Brand: brand,
        Model: model,
        RAM:   ram,
        Price: price,
    }
}

func main() {
    laptop := NewLaptop("Dell", "XPS 13", 16, 120000)
    fmt.Println("Laptop:", laptop.Brand, laptop.Model, laptop.RAM, "GB", laptop.Price)
}

```

---

✅ These exercises cover:

- **Struct creation** ✅
- **Constructor/factory functions** ✅
- **Methods on struct** ✅
- Real-world entities: Employee, BankAccount, Car, Book, Student, Product, Order, Movie, Customer, Laptop ✅