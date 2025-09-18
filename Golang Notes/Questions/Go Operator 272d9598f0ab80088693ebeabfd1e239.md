# Go Operator

## 🔹 10 Real-World Practice Questions on Go Operators

### 1. **Discount Calculation (Arithmetic Operator)**

A store offers a flat `20%` discount.

- Use arithmetic operators to calculate the final price after discount.

---

### 2. **Eligibility Check (Relational Operator)**

A movie requires viewers to be **at least 13 years old**.

- Use relational operators to check if a given age is eligible.

---

### 3. **Login Validation (Logical Operator)**

You need both a valid `username` and `password`.

- Use logical operators (`&&`, `||`) to print if login is successful.

---

### 4. **Bank Balance Update (Assignment Operator)**

A bank account starts with `1000`.

- Deposit `500`, withdraw `200` using compound assignment operators (`+=`, `=`).

---

### 5. **Voting System (Relational + Logical)**

A citizen can vote if they are **≥18 years old AND a registered voter**.

- Write a check using relational and logical operators.

---

### 6. **Like Counter (Increment Operator)**

A social media app shows a like counter starting at `0`.

- Simulate `++` when users like a post.

---

### 7. **Task Completion (Decrement Operator)**

A worker has `5` tasks.

- Use decrement operator (`-`) until all tasks are done.

---

### 8. **Even or Odd Check (Modulo Operator)**

Take an integer and use `%` to determine if it’s **even or odd**.

---

### 9. **Permission Flags (Bitwise AND/OR)**

You have permission flags:

- `Read = 4`, `Write = 2`, `Execute = 1`.
- A user has `5` (Read + Execute).
- Use bitwise operators to check if the user has `Write` permission.

---

### 10. **Shift Operator in Data Packets**

You receive data in bits.

- Use `<<` (left shift) to multiply by 2, and `>>` (right shift) to divide by 2.
- Example: A packet size is `8`. Show how shifting affects it.

---

### ✅ 1. Discount Calculation (Arithmetic Operator)

```go
package main

import "fmt"

func main() {
    var price float64 = 1000
    discount := 20.0

    finalPrice := price - (price * discount / 100)
    fmt.Printf("Final Price after %.0f%% discount: %.2f\n", discount, finalPrice)
}

```

---

### ✅ 2. Eligibility Check (Relational Operator)

```go
package main

import "fmt"

func main() {
    var age int = 12

    if age >= 13 {
        fmt.Println("Eligible to watch the movie")
    } else {
        fmt.Println("Not eligible to watch the movie")
    }
}

```

---

### ✅ 3. Login Validation (Logical Operator)

```go
package main

import "fmt"

func main() {
    username := "admin"
    password := "12345"

    if username == "admin" && password == "12345" {
        fmt.Println("Login Successful")
    } else {
        fmt.Println("Login Failed")
    }
}

```

---

### ✅ 4. Bank Balance Update (Assignment Operator)

```go
package main

import "fmt"

func main() {
    balance := 1000

    balance += 500  // deposit
    balance -= 200  // withdrawal

    fmt.Println("Final Balance:", balance)
}

```

---

### ✅ 5. Voting System (Relational + Logical)

```go
package main

import "fmt"

func main() {
    age := 20
    registered := true

    if age >= 18 && registered {
        fmt.Println("Eligible to vote")
    } else {
        fmt.Println("Not eligible to vote")
    }
}

```

---

### ✅ 6. Like Counter (Increment Operator)

```go
package main

import "fmt"

func main() {
    likes := 0

    likes++ // user 1 liked
    likes++ // user 2 liked
    likes++ // user 3 liked

    fmt.Println("Total Likes:", likes)
}

```

---

### ✅ 7. Task Completion (Decrement Operator)

```go
package main

import "fmt"

func main() {
    tasks := 5

    for tasks > 0 {
        fmt.Println("Task completed. Remaining:", tasks-1)
        tasks-- // decrement
    }
}

```

---

### ✅ 8. Even or Odd Check (Modulo Operator)

```go
package main

import "fmt"

func main() {
    num := 7

    if num%2 == 0 {
        fmt.Println(num, "is Even")
    } else {
        fmt.Println(num, "is Odd")
    }
}

```

---

### ✅ 9. Permission Flags (Bitwise AND/OR)

```go
package main

import "fmt"

func main() {
    const Read = 4
    const Write = 2
    const Execute = 1

    userPermission := 5 // Read + Execute

    if userPermission&Write != 0 {
        fmt.Println("User has Write permission")
    } else {
        fmt.Println("User does NOT have Write permission")
    }
}

```

---

### ✅ 10. Shift Operator in Data Packets

```go
package main

import "fmt"

func main() {
    packetSize := 8

    doubled := packetSize << 1 // multiply by 2
    halved := packetSize >> 1  // divide by 2

    fmt.Println("Original:", packetSize)
    fmt.Println("Doubled (<< 1):", doubled)
    fmt.Println("Halved (>> 1):", halved)
}

```

---

⚡ Now you have **complete working Go programs** for all operator types:

- Arithmetic
- Relational
- Logical
- Assignment
- Increment / Decrement
- Modulo
- Bitwise
- Shift