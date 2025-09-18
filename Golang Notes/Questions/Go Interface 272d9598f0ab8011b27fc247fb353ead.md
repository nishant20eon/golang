# Go Interface

## 🔹 10 Real-World Practice Questions on Go Interfaces

### 1. **Shape Interface**

- Create an interface `Shape` with method `Area() float64`.
- Implement `Rectangle` and `Circle` structs with `Area()` methods.
- Write a function that takes `Shape` and prints its area.

---

### 2. **Payment Interface**

- Create interface `Payment` with method `Pay(amount float64)`.
- Implement `CreditCard` and `PayPal` structs that implement `Pay`.
- Demonstrate polymorphic behavior by paying using both.

---

### 3. **Vehicle Interface**

- Interface `Vehicle` with method `Drive()`.
- Implement `Car` and `Bike` structs with `Drive()` methods.
- Iterate over a slice of `Vehicle` and call `Drive()` on each.

---

### 4. **Notifier Interface**

- Interface `Notifier` with method `Notify(message string)`.
- Implement `EmailNotifier` and `SMSNotifier`.
- Send a message using both notifiers.

---

### 5. **Employee Bonus Interface**

- Interface `BonusCalculator` with method `CalculateBonus() float64`.
- Implement `Manager` and `Developer` structs.
- Print bonus for a slice of `BonusCalculator`.

---

### 6. **Logger Interface**

- Interface `Logger` with method `Log(message string)`.
- Implement `ConsoleLogger` and `FileLogger` (simulate file log with fmt).
- Log the same message using both loggers.

---

### 7. **Drawable Interface**

- Interface `Drawable` with method `Draw()`.
- Implement `Circle`, `Rectangle`, `Triangle`.
- Use a function to draw all shapes in a slice of `Drawable`.

---

### 8. **Taxable Interface**

- Interface `Taxable` with method `CalculateTax() float64`.
- Implement `Product` and `Service` structs.
- Calculate tax for a slice of `Taxable` items.

---

### 9. **Animal Sound Interface**

- Interface `Animal` with method `Sound() string`.
- Implement `Dog` and `Cat` structs.
- Use a slice of `Animal` to print all sounds.

---

### 🔟 **Shape Comparison Interface**

- Interface `Shape` with method `Area() float64`.
- Implement `Rectangle` and `Circle`.
- Write a function that takes two `Shape` and prints which one has a larger area.

---

⚡ These exercises cover:

- **Defining interfaces and methods** ✅
- **Polymorphism**: calling methods on different types through the same interface ✅
- **Slices of interface types** ✅
- **Real-world entities**: shapes, payments, vehicles, notifiers, employees, animals ✅

---

## 1️⃣ Shape Interface

```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Rectangle struct {
    Length, Width float64
}

func (r Rectangle) Area() float64 {
    return r.Length * r.Width
}

type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func PrintArea(s Shape) {
    fmt.Println("Area:", s.Area())
}

func main() {
    r := Rectangle{10, 5}
    c := Circle{7}
    PrintArea(r)
    PrintArea(c)
}

```

---

## 2️⃣ Payment Interface

```go
package main

import "fmt"

type Payment interface {
    Pay(amount float64)
}

type CreditCard struct {
    Owner string
}

func (c CreditCard) Pay(amount float64) {
    fmt.Printf("%s paid %.2f using Credit Card\n", c.Owner, amount)
}

type PayPal struct {
    Email string
}

func (p PayPal) Pay(amount float64) {
    fmt.Printf("%s paid %.2f using PayPal\n", p.Email, amount)
}

func main() {
    var p Payment

    p = CreditCard{"Nishant"}
    p.Pay(1000)

    p = PayPal{"nishant@example.com"}
    p.Pay(2000)
}

```

---

## 3️⃣ Vehicle Interface

```go
package main

import "fmt"

type Vehicle interface {
    Drive()
}

type Car struct{}
func (c Car) Drive() { fmt.Println("Car is driving") }

type Bike struct{}
func (b Bike) Drive() { fmt.Println("Bike is driving") }

func main() {
    vehicles := []Vehicle{Car{}, Bike{}}
    for _, v := range vehicles {
        v.Drive()
    }
}

```

---

## 4️⃣ Notifier Interface

```go
package main

import "fmt"

type Notifier interface {
    Notify(message string)
}

type EmailNotifier struct{}
func (e EmailNotifier) Notify(message string) {
    fmt.Println("Email:", message)
}

type SMSNotifier struct{}
func (s SMSNotifier) Notify(message string) {
    fmt.Println("SMS:", message)
}

func main() {
    notifiers := []Notifier{EmailNotifier{}, SMSNotifier{}}
    for _, n := range notifiers {
        n.Notify("Hello, Go Interfaces!")
    }
}

```

---

## 5️⃣ Employee Bonus Interface

```go
package main

import "fmt"

type BonusCalculator interface {
    CalculateBonus() float64
}

type Manager struct {
    Salary float64
}
func (m Manager) CalculateBonus() float64 { return m.Salary * 0.2 }

type Developer struct {
    Salary float64
}
func (d Developer) CalculateBonus() float64 { return d.Salary * 0.1 }

func main() {
    employees := []BonusCalculator{Manager{50000}, Developer{40000}}
    for _, e := range employees {
        fmt.Println("Bonus:", e.CalculateBonus())
    }
}

```

---

## 6️⃣ Logger Interface

```go
package main

import "fmt"

type Logger interface {
    Log(message string)
}

type ConsoleLogger struct{}
func (c ConsoleLogger) Log(message string) { fmt.Println("Console:", message) }

type FileLogger struct{}
func (f FileLogger) Log(message string) { fmt.Println("File:", message) } // simulate file log

func main() {
    loggers := []Logger{ConsoleLogger{}, FileLogger{}}
    for _, l := range loggers {
        l.Log("Logging message")
    }
}

```

---

## 7️⃣ Drawable Interface

```go
package main

import "fmt"

type Drawable interface {
    Draw()
}

type Circle struct{}
func (c Circle) Draw() { fmt.Println("Drawing Circle") }

type Rectangle struct{}
func (r Rectangle) Draw() { fmt.Println("Drawing Rectangle") }

type Triangle struct{}
func (t Triangle) Draw() { fmt.Println("Drawing Triangle") }

func main() {
    shapes := []Drawable{Circle{}, Rectangle{}, Triangle{}}
    for _, s := range shapes {
        s.Draw()
    }
}

```

---

## 8️⃣ Taxable Interface

```go
package main

import "fmt"

type Taxable interface {
    CalculateTax() float64
}

type Product struct {
    Price float64
}
func (p Product) CalculateTax() float64 { return p.Price * 0.1 }

type Service struct {
    Amount float64
}
func (s Service) CalculateTax() float64 { return s.Amount * 0.15 }

func main() {
    items := []Taxable{Product{100}, Service{200}}
    for _, item := range items {
        fmt.Println("Tax:", item.CalculateTax())
    }
}

```

---

## 9️⃣ Animal Sound Interface

```go
package main

import "fmt"

type Animal interface {
    Sound() string
}

type Dog struct{}
func (d Dog) Sound() string { return "Woof" }

type Cat struct{}
func (c Cat) Sound() string { return "Meow" }

func main() {
    animals := []Animal{Dog{}, Cat{}}
    for _, a := range animals {
        fmt.Println(a.Sound())
    }
}

```

---

## 🔟 Shape Comparison Interface

```go
package main

import (
    "fmt"
    "math"
)

type Shape interface {
    Area() float64
}

type Rectangle struct{ Length, Width float64 }
func (r Rectangle) Area() float64 { return r.Length * r.Width }

type Circle struct{ Radius float64 }
func (c Circle) Area() float64 { return math.Pi * c.Radius * c.Radius }

func CompareAreas(s1, s2 Shape) {
    if s1.Area() > s2.Area() {
        fmt.Println("First shape has larger area")
    } else {
        fmt.Println("Second shape has larger area")
    }
}

func main() {
    r := Rectangle{10, 5}
    c := Circle{7}
    CompareAreas(r, c)
}

```

---

✅ These exercises cover:

- **Defining interfaces & methods** ✅
- **Polymorphism**: multiple types implementing same interface ✅
- **Slices of interface types** ✅
- **Real-world scenarios**: shapes, payments, vehicles, notifiers, animals, employees ✅