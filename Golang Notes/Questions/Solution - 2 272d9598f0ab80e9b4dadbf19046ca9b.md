# Solution - 2

## 📂 Project Structure

```
go-functions-practice/
│
├── go.mod
├── main.go
├── mathutils/
│   └── mathutils.go
├── greet/
│   └── greet.go
├── converter/
│   └── converter.go
├── payment/
│   └── payment.go
├── auth/
│   └── auth.go
├── bank/
│   └── bank.go
├── student/
│   └── student.go
├── order/
│   └── order.go
├── car/
│   └── car.go
└── report/
    └── report.go

```

---

## 1️⃣ `go.mod`

```go
module go-functions-practice

go 1.21

```

---

## 2️⃣ `mathutils/mathutils.go`

```go
package mathutils

func Square(n int) int {
    return n * n
}

```

---

## 3️⃣ `greet/greet.go`

```go
package greet

func Hello(name string) string {
    return "Hello " + name
}

```

---

## 4️⃣ `converter/converter.go`

```go
package converter

func CelsiusToFahrenheit(c float64) float64 {
    return (c*9/5 + 32)
}

```

---

## 5️⃣ `payment/payment.go`

```go
package payment

func Discount(price float64, percent float64) float64 {
    return price - (price * percent / 100)
}

```

---

## 6️⃣ `auth/auth.go`

```go
package auth

import "errors"

func ValidateUser(username, password string) (bool, error) {
    if username == "admin" && password == "1234" {
        return true, nil
    }
    return false, errors.New("invalid credentials")
}

```

---

## 7️⃣ `bank/bank.go`

```go
package bank

type Account struct {
    Balance float64
}

// Value Receiver
func (a Account) ShowBalance() float64 {
    return a.Balance
}

// Pointer Receiver
func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}

```

---

## 8️⃣ `student/student.go`

```go
package student

func Average(marks []int) float64 {
    total := 0
    for _, m := range marks {
        total += m
    }
    return float64(total) / float64(len(marks))
}

```

---

## 9️⃣ `order/order.go`

```go
package order

func CalculateTotal(prices []float64) (float64, int) {
    total := 0.0
    for _, p := range prices {
        total += p
    }
    return total, len(prices)
}

```

---

## 🔟 `car/car.go`

```go
package car

type Car struct {
    Brand   string
    Mileage int
}

func (c *Car) AddMileage(miles int) {
    c.Mileage += miles
}

```

---

## 1️⃣1️⃣ `report/report.go`

```go
package report

import (
    "fmt"
    "go-functions-practice/car"
)

func MileageReport(c *car.Car) {
    fmt.Printf("Car %s has Mileage: %d\n", c.Brand, c.Mileage)
}

```

---

## 1️⃣2️⃣ `main.go`

```go
package main

import (
    "fmt"

    "go-functions-practice/auth"
    "go-functions-practice/bank"
    "go-functions-practice/car"
    "go-functions-practice/converter"
    "go-functions-practice/greet"
    "go-functions-practice/mathutils"
    "go-functions-practice/order"
    "go-functions-practice/payment"
    "go-functions-practice/report"
    "go-functions-practice/student"
)

func main() {
    // 1. Cube using mathutils
    fmt.Println("Cube of 3:", 3*mathutils.Square(3))

    // 2. Welcome User using greet
    fmt.Println(greet.Hello("Nishant") + ", Welcome to our system!")

    // 3. Weather Report using converter
    fmt.Printf("Today's temperature: %.2f°F\n", converter.CelsiusToFahrenheit(30))

    // 4. Final Amount using payment
    fmt.Printf("Final Price after discount: %.2f\n", payment.Discount(1000, 10))

    // 5. Login using auth
    valid, _ := auth.ValidateUser("admin", "1234")
    if valid {
        fmt.Println("Login Success")
    } else {
        fmt.Println("Login Failed")
    }

    // 6. Bank Account Value Receiver
    acc1 := bank.Account{Balance: 5000}
    fmt.Println("Balance:", acc1.ShowBalance())

    // 7. Bank Account Pointer Receiver
    acc2 := &bank.Account{Balance: 5000}
    acc2.Deposit(2000)
    fmt.Println("Updated Balance after Deposit:", acc2.Balance)

    // 8. Student Average & Result
    marks := []int{50, 60, 70}
    avg := student.Average(marks)
    if avg >= 40 {
        fmt.Println("Pass")
    } else {
        fmt.Println("Fail")
    }

    // 9. Order Summary using order
    prices := []float64{100, 200, 150}
    total, count := order.CalculateTotal(prices)
    fmt.Printf("Total Items: %d, Total Price: %.2f\n", count, total)

    // 10. Car Travel using car & report
    myCar := &car.Car{"Tesla", 12000}
    myCar.AddMileage(500)
    report.MileageReport(myCar)
}

```

---

✅ This **full project template** allows you to:

- Practice **cross-package function calls**
- Use **value & pointer receivers**
- Handle **single/multi param & return types**
- Run everything from `main.go`