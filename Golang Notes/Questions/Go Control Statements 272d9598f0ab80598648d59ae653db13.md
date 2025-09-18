# Go Control Statements

## 🔹 10 Real-World Practice Questions on Go Control Statements

### 1. **Login Attempts (if-else + loop)**

Allow a user a maximum of **3 login attempts**.

- If credentials are wrong 3 times, print `"Account Locked"`.
- Otherwise, `"Login Successful"`.

---

### 2. **Traffic Light System (switch-case)**

Simulate a traffic signal:

- `"Red"` → Stop, `"Yellow"` → Get Ready, `"Green"` → Go.
- Use `switch` to print actions.

---

### 3. **ATM Withdrawal (if-else)**

Check if the withdrawal amount is less than or equal to balance.

- If yes, deduct it.
- Else, print `"Insufficient Balance"`.

---

### 4. **Even Numbers Printer (for loop + continue)**

Print numbers from 1–20 but **skip odd numbers** using `continue`.

---

### 5. **First Prime Number Found (for loop + break)**

Loop through numbers from 10–30.

- Find the first prime number.
- Stop the loop using `break`.

---

### 6. **Simple Menu (switch + loop)**

Create a console menu:

1. Show Balance
2. Deposit
3. Withdraw
4. Exit
- Keep looping until user chooses Exit.

---

### 7. **Student Pass/Fail (nested if)**

Check if a student has:

- `marks >= 40` → Pass,
- else → Fail.
    
    If passed, further check:
    
- `marks >= 75` → Distinction.

---

### 8. **Password Retry (for + break)**

Ask for a password in a loop.

- If correct, print `"Access Granted"` and exit loop.
- If wrong, keep asking until correct.

---

### 9. **Skip Weekends (for + continue)**

Print days of the week (1–7).

- Skip day `6` (Saturday) and `7` (Sunday).

---

### 10. **Goto Statement Example**

Simulate an emergency exit in a program using `goto` to jump to a label when an error occurs.

### ✅ 1. Login Attempts (if-else + loop)

```go
package main

import "fmt"

func main() {
    correctUser, correctPass := "admin", "1234"
    attempts := 3

    for i := 1; i <= attempts; i++ {
        var user, pass string
        fmt.Printf("Attempt %d - Enter username: ", i)
        fmt.Scanln(&user)
        fmt.Print("Enter password: ")
        fmt.Scanln(&pass)

        if user == correctUser && pass == correctPass {
            fmt.Println("Login Successful ✅")
            return
        } else {
            fmt.Println("Wrong credentials ❌")
        }
    }
    fmt.Println("Account Locked 🔒")
}

```

---

### ✅ 2. Traffic Light System (switch-case)

```go
package main

import "fmt"

func main() {
    signal := "Yellow"

    switch signal {
    case "Red":
        fmt.Println("Stop 🚦")
    case "Yellow":
        fmt.Println("Get Ready ⚠️")
    case "Green":
        fmt.Println("Go ✅")
    default:
        fmt.Println("Invalid signal")
    }
}

```

---

### ✅ 3. ATM Withdrawal (if-else)

```go
package main

import "fmt"

func main() {
    balance := 1000
    withdrawal := 1200

    if withdrawal <= balance {
        balance -= withdrawal
        fmt.Println("Withdrawal Successful. Remaining Balance:", balance)
    } else {
        fmt.Println("Insufficient Balance ❌")
    }
}

```

---

### ✅ 4. Even Numbers Printer (for loop + continue)

```go
package main

import "fmt"

func main() {
    for i := 1; i <= 20; i++ {
        if i%2 != 0 {
            continue // skip odd numbers
        }
        fmt.Print(i, " ")
    }
}

```

---

### ✅ 5. First Prime Number Found (for loop + break)

```go
package main

import "fmt"

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func main() {
    for i := 10; i <= 30; i++ {
        if isPrime(i) {
            fmt.Println("First Prime Number:", i)
            break
        }
    }
}

```

---

### ✅ 6. Simple Menu (switch + loop)

```go
package main

import "fmt"

func main() {
    balance := 1000
    for {
        fmt.Println("\n--- Menu ---")
        fmt.Println("1. Show Balance")
        fmt.Println("2. Deposit")
        fmt.Println("3. Withdraw")
        fmt.Println("4. Exit")

        var choice int
        fmt.Print("Enter choice: ")
        fmt.Scanln(&choice)

        switch choice {
        case 1:
            fmt.Println("Balance:", balance)
        case 2:
            var amt int
            fmt.Print("Enter deposit amount: ")
            fmt.Scanln(&amt)
            balance += amt
            fmt.Println("Deposited:", amt)
        case 3:
            var amt int
            fmt.Print("Enter withdrawal amount: ")
            fmt.Scanln(&amt)
            if amt <= balance {
                balance -= amt
                fmt.Println("Withdrawn:", amt)
            } else {
                fmt.Println("Insufficient Balance ❌")
            }
        case 4:
            fmt.Println("Exiting... 👋")
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}

```

---

### ✅ 7. Student Pass/Fail (nested if)

```go
package main

import "fmt"

func main() {
    marks := 80

    if marks >= 40 {
        fmt.Println("Pass ✅")
        if marks >= 75 {
            fmt.Println("Distinction 🎉")
        }
    } else {
        fmt.Println("Fail ❌")
    }
}

```

---

### ✅ 8. Password Retry (for + break)

```go
package main

import "fmt"

func main() {
    correctPassword := "golang"

    for {
        var pass string
        fmt.Print("Enter password: ")
        fmt.Scanln(&pass)

        if pass == correctPassword {
            fmt.Println("Access Granted ✅")
            break
        } else {
            fmt.Println("Wrong password ❌, try again...")
        }
    }
}

```

---

### ✅ 9. Skip Weekends (for + continue)

```go
package main

import "fmt"

func main() {
    for day := 1; day <= 7; day++ {
        if day == 6 || day == 7 {
            continue // skip Saturday & Sunday
        }
        fmt.Println("Working Day:", day)
    }
}

```

---

### ✅ 10. Goto Statement Example

```go
package main

import "fmt"

func main() {
    var errorOccurred = true

    if errorOccurred {
        goto emergencyExit
    }

    fmt.Println("Program running normally")

emergencyExit:
    fmt.Println("Emergency Exit Triggered ⚠️")
}

```

---

✅ Now you have **full working Go examples** for:

- `if-else`, `nested if`
- `switch`
- `for loop` with `break` / `continue`
- `goto`