# Go Function

## 🔹 15 Real-World Practice Questions on Functions in Go

### **Single Parameter**

1. **Square of a Number**
    
    Write a function `square(num int) int` that returns the square of a number.
    
2. **Greeting Message**
    
    Create a function `greet(name string)` that prints `"Hello <name>"`.
    
3. **Check Even/Odd**
    
    Write a function `isEven(n int) bool` that returns `true` if number is even, else `false`.
    

---

### **Multiple Parameters**

1. **Add Two Numbers**
    
    Create a function `add(a int, b int) int` that returns the sum.
    
2. **Area of Rectangle**
    
    Write a function `area(length, width float64) float64` that returns the area.
    
3. **BMI Calculator**
    
    Function `bmi(weight, height float64) float64` should return BMI = weight / (height * height).
    

---

### **Single Return Type**

1. **Factorial**
    
    Write `factorial(n int) int` that calculates factorial using recursion.
    
2. **Reverse String**
    
    Write `reverse(s string) string` that reverses a string.
    

---

### **Multiple Return Types**

1. **Divide Numbers**
    
    Write `divide(a, b float64) (float64, error)` that returns quotient and error if `b==0`.
    
2. **Min and Max**
    
    Write `minMax(nums []int) (int, int)` that returns smallest and largest number in a slice.
    

---

### **Normal Receiver (Value Receiver)**

1. **Book Info**
    
    Define a struct `Book{Title, Author string}`.
    
    Write a method `(b Book) printInfo()` that prints book details.
    
2. **Circle Area**
    
    Define `Circle{Radius float64}`.
    
    Method `(c Circle) area() float64` should return circle’s area.
    

---

### **Pointer Receiver**

1. **Bank Account Deposit**
    
    Define `Account{Balance float64}`.
    
    Write method `(a *Account) deposit(amount float64)` that updates balance.
    
2. **Update Student Marks**
    
    Struct `Student{Name string, Marks int}`.
    
    Write method `(s *Student) updateMarks(newMarks int)` to change marks.
    
3. **Car Mileage Update**
    
    Struct `Car{Brand string, Mileage int}`.
    
    Write method `(c *Car) addMileage(miles int)` that increases mileage.
    

---

## ✅ 1. Square of a Number (Single Param, Single Return)

```go
package main

import "fmt"

func square(num int) int {
    return num * num
}

func main() {
    fmt.Println("Square:", square(5))
}

```

---

## ✅ 2. Greeting Message (Single Param, No Return)

```go
package main

import "fmt"

func greet(name string) {
    fmt.Println("Hello", name)
}

func main() {
    greet("Nishant")
}

```

---

## ✅ 3. Check Even/Odd (Single Param, Bool Return)

```go
package main

import "fmt"

func isEven(n int) bool {
    return n%2 == 0
}

func main() {
    fmt.Println("Is 10 Even?", isEven(10))
    fmt.Println("Is 7 Even?", isEven(7))
}

```

---

## ✅ 4. Add Two Numbers (Multi Param, Single Return)

```go
package main

import "fmt"

func add(a int, b int) int {
    return a + b
}

func main() {
    fmt.Println("Sum:", add(5, 7))
}

```

---

## ✅ 5. Area of Rectangle (Multi Param, Single Return)

```go
package main

import "fmt"

func area(length, width float64) float64 {
    return length * width
}

func main() {
    fmt.Println("Area:", area(10, 5))
}

```

---

## ✅ 6. BMI Calculator (Multi Param, Single Return)

```go
package main

import "fmt"

func bmi(weight, height float64) float64 {
    return weight / (height * height)
}

func main() {
    fmt.Printf("BMI: %.2f\n", bmi(70, 1.75))
}

```

---

## ✅ 7. Factorial (Recursion, Single Return)

```go
package main

import "fmt"

func factorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorial(n-1)
}

func main() {
    fmt.Println("Factorial of 5:", factorial(5))
}

```

---

## ✅ 8. Reverse String (Single Param, Single Return)

```go
package main

import "fmt"

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    fmt.Println("Reverse of GoLang:", reverse("GoLang"))
}

```

---

## ✅ 9. Divide Numbers (Multi Return: Value + Error)

```go
package main

import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 2)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Result:", result)
    }
}

```

---

## ✅ 10. Min and Max in Slice (Multi Return)

```go
package main

import "fmt"

func minMax(nums []int) (int, int) {
    min, max := nums[0], nums[0]
    for _, v := range nums {
        if v < min {
            min = v
        }
        if v > max {
            max = v
        }
    }
    return min, max
}

func main() {
    numbers := []int{5, 8, 1, 9, 3}
    min, max := minMax(numbers)
    fmt.Println("Min:", min, "Max:", max)
}

```

---

## ✅ 11. Book Info (Value Receiver)

```go
package main

import "fmt"

type Book struct {
    Title  string
    Author string
}

func (b Book) printInfo() {
    fmt.Println("Book:", b.Title, "by", b.Author)
}

func main() {
    b := Book{"Go Programming", "John Doe"}
    b.printInfo()
}

```

---

## ✅ 12. Circle Area (Value Receiver)

```go
package main

import (
    "fmt"
    "math"
)

type Circle struct {
    Radius float64
}

func (c Circle) area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func main() {
    c := Circle{Radius: 7}
    fmt.Printf("Circle Area: %.2f\n", c.area())
}

```

---

## ✅ 13. Bank Account Deposit (Pointer Receiver)

```go
package main

import "fmt"

type Account struct {
    Balance float64
}

func (a *Account) deposit(amount float64) {
    a.Balance += amount
}

func main() {
    acc := Account{Balance: 1000}
    acc.deposit(500)
    fmt.Println("Updated Balance:", acc.Balance)
}

```

---

## ✅ 14. Update Student Marks (Pointer Receiver)

```go
package main

import "fmt"

type Student struct {
    Name  string
    Marks int
}

func (s *Student) updateMarks(newMarks int) {
    s.Marks = newMarks
}

func main() {
    stu := Student{"Nishant", 70}
    stu.updateMarks(90)
    fmt.Println("Updated Student:", stu.Name, "Marks:", stu.Marks)
}

```

---

## ✅ 15. Car Mileage Update (Pointer Receiver)

```go
package main

import "fmt"

type Car struct {
    Brand   string
    Mileage int
}

func (c *Car) addMileage(miles int) {
    c.Mileage += miles
}

func main() {
    car := Car{"Tesla", 12000}
    car.addMileage(500)
    fmt.Println("Updated Car:", car.Brand, "Mileage:", car.Mileage)
}

```

---

⚡ These 15 programs cover **all kinds of functions in Go**:

- Single param ✅
- Multi param ✅
- Single return ✅
- Multi return ✅
- Value receiver methods ✅
- Pointer receiver methods ✅