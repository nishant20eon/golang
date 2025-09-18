# Go Datatype

### 🔹 1. User Age Validation

You are building a registration form. Write a Go program that:

- Declares a variable `age` of type `int`.
- Checks if the age is above 18.
- Prints `"Eligible"` or `"Not Eligible"`.

---

### 🔹 2. Temperature Conversion

Declare a constant `freezingCelsius = 0`.

- Take a variable `tempFahrenheit` (float64).
- Convert it into Celsius using formula: `(F - 32) * 5/9`.
- Print whether it is below or above freezing point.

---

### 🔹 3. Employee Salary with Bonus

You have a base salary stored in a variable and a constant `BONUS_PERCENTAGE = 10`.

- Write a program to calculate total salary after applying the bonus.

---

### 🔹 4. E-commerce Order IDs

You need to manage a list of order IDs.

- Use a slice of type `int` to store multiple order IDs.
- Print the total number of orders using `len(slice)`.

---

### 🔹 5. Bank Balance Precision

- Declare a variable `balance` of type `float64`.
- Add deposit and subtract withdrawal values.
- Print the final balance with **2 decimal places**.

---

### 🔹 6. Student Grade Evaluation

- Declare a `marks` variable of type `int`.
- If marks ≥ 90 → print `"A"`, 70–89 → `"B"`, else `"C"`.

---

### 🔹 7. Chat Message Character Count

You want to limit messages in a chat app.

- Declare a variable `message` of type `string`.
- Print the number of characters using `len(message)`.
- If more than 200, print `"Message too long"`.

---

### 🔹 8. IoT Sensor Readings

- Store temperature readings from sensors in a slice of type `float64`.
- Print average temperature using loop and length.

---

### 🔹 9. Currency Conversion

You have a constant `USD_TO_INR = 83.2`.

- Take a variable `usdAmount` (float64).
- Convert to INR and print the result.

---

### 🔹 10. Shopping Cart Prices

- Use a slice of type `float64` to store item prices.
- Add a constant `TAX_RATE = 0.18`.
- Calculate total price with tax and print.

### ✅ 1. User Age Validation

```go
package main

import "fmt"

func main() {
    var age int = 20
    if age >= 18 {
        fmt.Println("Eligible")
    } else {
        fmt.Println("Not Eligible")
    }
}

```

---

### ✅ 2. Temperature Conversion

```go
package main

import "fmt"

func main() {
    const freezingCelsius = 0
    var tempFahrenheit float64 = 45.0

    celsius := (tempFahrenheit - 32) * 5 / 9
    fmt.Println("Temperature in Celsius:", celsius)

    if celsius <= float64(freezingCelsius) {
        fmt.Println("Below freezing point")
    } else {
        fmt.Println("Above freezing point")
    }
}

```

---

### ✅ 3. Employee Salary with Bonus

```go
package main

import "fmt"

func main() {
    const BONUS_PERCENTAGE = 10
    var baseSalary float64 = 50000

    totalSalary := baseSalary + (baseSalary * float64(BONUS_PERCENTAGE) / 100)
    fmt.Println("Total Salary:", totalSalary)
}

```

---

### ✅ 4. E-commerce Order IDs

```go
package main

import "fmt"

func main() {
    orders := []int{101, 102, 103, 104, 105}
    fmt.Println("Order IDs:", orders)
    fmt.Println("Total number of orders:", len(orders))
}

```

---

### ✅ 5. Bank Balance Precision

```go
package main

import "fmt"

func main() {
    var balance float64 = 1000.75
    deposit := 500.25
    withdrawal := 200.50

    balance = balance + deposit - withdrawal
    fmt.Printf("Final Balance: %.2f\n", balance)
}

```

---

### ✅ 6. Student Grade Evaluation

```go
package main

import "fmt"

func main() {
    var marks int = 85

    if marks >= 90 {
        fmt.Println("Grade: A")
    } else if marks >= 70 {
        fmt.Println("Grade: B")
    } else {
        fmt.Println("Grade: C")
    }
}

```

---

### ✅ 7. Chat Message Character Count

```go
package main

import "fmt"

func main() {
    var message string = "This is a test chat message."

    fmt.Println("Message length:", len(message))

    if len(message) > 200 {
        fmt.Println("Message too long")
    } else {
        fmt.Println("Message is within limit")
    }
}

```

---

### ✅ 8. IoT Sensor Readings

```go
package main

import "fmt"

func main() {
    readings := []float64{22.5, 23.0, 21.8, 22.9}
    sum := 0.0

    for _, v := range readings {
        sum += v
    }

    avg := sum / float64(len(readings))
    fmt.Println("Average Temperature:", avg)
}

```

---

### ✅ 9. Currency Conversion

```go
package main

import "fmt"

func main() {
    const USD_TO_INR = 83.2
    var usdAmount float64 = 100

    inr := usdAmount * USD_TO_INR
    fmt.Printf("USD %.2f = INR %.2f\n", usdAmount, inr)
}

```

---

### ✅ 10. Shopping Cart Prices

```go
package main

import "fmt"

func main() {
    cart := []float64{100.50, 200.75, 50.25}
    const TAX_RATE = 0.18

    total := 0.0
    for _, price := range cart {
        total += price
    }

    totalWithTax := total + (total * TAX_RATE)
    fmt.Printf("Total Price with Tax: %.2f\n", totalWithTax)
}

```