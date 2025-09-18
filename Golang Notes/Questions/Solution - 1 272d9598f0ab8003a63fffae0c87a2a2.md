# Solution - 1

## 📂 Suggested File Structure

```
go-functions-practice/
│
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

## 1️⃣ Package: `mathutils` – Square

```go
package mathutils

func Square(n int) int {
    return n * n
}

```

### Main function calling Cube

```go
package main

import (
    "fmt"
    "go-functions-practice/mathutils"
)

func Cube(n int) int {
    return n * mathutils.Square(n)
}

func main() {
    fmt.Println("Cube of 3:", Cube(3))
}

```

---

## 2️⃣ Package: `greet`

```go
package greet

func Hello(name string) string {
    return "Hello " + name
}

```

### Main function calling WelcomeUser

```go
package main

import (
    "fmt"
    "go-functions-practice/greet"
)

func WelcomeUser(name string) string {
    return greet.Hello(name) + ", Welcome to our system!"
}

func main() {
    fmt.Println(WelcomeUser("Nishant"))
}

```

---

## 3️⃣ Package: `converter`

```go
package converter

func CelsiusToFahrenheit(c float64) float64 {
    return (c*9/5 + 32)
}

```

### Main function calling WeatherReport

```go
package main

import (
    "fmt"
    "go-functions-practice/converter"
)

func WeatherReport(celsius float64) string {
    fahrenheit := converter.CelsiusToFahrenheit(celsius)
    return fmt.Sprintf("Today's temperature: %.2f°F", fahrenheit)
}

func main() {
    fmt.Println(WeatherReport(30))
}

```

---

## 4️⃣ Package: `payment`

```go
package payment

func Discount(price float64, percent float64) float64 {
    return price - (price * percent / 100)
}

```

### Main function calling FinalAmount

```go
package main

import (
    "fmt"
    "go-functions-practice/payment"
)

func FinalAmount(price float64, discount float64) float64 {
    return payment.Discount(price, discount)
}

func main() {
    fmt.Printf("Final Price: %.2f\n", FinalAmount(1000, 10))
}

```

---

## 5️⃣ Package: `auth`

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

### Main function calling Login

```go
package main

import (
    "fmt"
    "go-functions-practice/auth"
)

func Login(username, password string) string {
    valid, _ := auth.ValidateUser(username, password)
    if valid {
        return "Login Success"
    }
    return "Login Failed"
}

func main() {
    fmt.Println(Login("admin", "1234"))
    fmt.Println(Login("user", "pass"))
}

```

---

## 6️⃣ Package: `bank` (Value Receiver)

```go
package bank

type Account struct {
    Balance float64
}

func (a Account) ShowBalance() float64 {
    return a.Balance
}

```

### Main function calling CheckBalance

```go
package main

import (
    "fmt"
    "go-functions-practice/bank"
)

func CheckBalance(a bank.Account) {
    fmt.Println("Balance:", a.ShowBalance())
}

func main() {
    acc := bank.Account{Balance: 5000}
    CheckBalance(acc)
}

```

---

## 7️⃣ Package: `bank` (Pointer Receiver)

```go
package bank

func (a *Account) Deposit(amount float64) {
    a.Balance += amount
}

```

### Main function calling SalaryCredit

```go
package main

import (
    "fmt"
    "go-functions-practice/bank"
)

func SalaryCredit(a *bank.Account, salary float64) {
    a.Deposit(salary)
}

func main() {
    acc := &bank.Account{Balance: 5000}
    SalaryCredit(acc, 2000)
    fmt.Println("Updated Balance:", acc.Balance)
}

```

---

## 8️⃣ Package: `student`

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

### Main function calling ResultStatus

```go
package main

import (
    "fmt"
    "go-functions-practice/student"
)

func ResultStatus(marks []int) string {
    avg := student.Average(marks)
    if avg >= 40 {
        return "Pass"
    }
    return "Fail"
}

func main() {
    marks := []int{50, 60, 70}
    fmt.Println(ResultStatus(marks))
}

```

---

## 9️⃣ Package: `order`

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

### Main function calling OrderSummary

```go
package main

import (
    "fmt"
    "go-functions-practice/order"
)

func OrderSummary(prices []float64) string {
    total, count := order.CalculateTotal(prices)
    return fmt.Sprintf("Total Items: %d, Total Price: %.2f", count, total)
}

func main() {
    prices := []float64{100, 200, 150}
    fmt.Println(OrderSummary(prices))
}

```

---

## 🔟 Packages: `car` & `report`

### `car/car.go`

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

### `report/report.go`

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

### Main function calling Travel

```go
package main

import (
    "go-functions-practice/car"
    "go-functions-practice/report"
)

func Travel(c *car.Car, miles int) {
    c.AddMileage(miles)
    report.MileageReport(c)
}

func main() {
    myCar := &car.Car{"Tesla", 12000}
    Travel(myCar, 500)
}

```

---

✅ These examples cover:

- **Single param & multi param**
- **Single return & multi return**
- **Value & pointer receivers**
- **Functions calling functions from different packages**