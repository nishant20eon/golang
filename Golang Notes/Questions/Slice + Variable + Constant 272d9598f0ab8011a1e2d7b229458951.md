# Slice + Variable + Constant

## 🔹 10 Real-World Practice Questions (Slice + Variable + Constant)

1. **Student Marks**
    - Declare a constant for number of subjects.
    - Use a slice of integers (marks) and calculate the average.
2. **Shopping Cart**
    - Create a slice of item prices.
    - Use a variable `discount` and calculate the total after applying discount.
3. **Employee List**
    - Create a slice of employee names.
    - Append a new employee name (variable input).
4. **Todo Tasks**
    - Define a constant for `MaxTasks`.
    - Add tasks into a slice until it reaches max size. Print all tasks.
5. **Temperature Records**
    - Store daily temperatures in a slice.
    - Find max and min temperature using variables.
6. **Bank Transactions**
    - Store amounts in a slice.
    - Use a constant transaction fee and deduct it from each transaction.
7. **Top Scorers**
    - Keep a slice of player scores.
    - Sort the slice and print top 3 scorers using constants.
8. **Library Books**
    - Maintain a slice of book titles.
    - Remove a book from the slice based on a variable index.
9. **Student Grades**
    - Use a slice of integers for marks.
    - Use a constant passing mark.
    - Count how many students passed/failed using variables.
10. **Online Orders**
- Maintain a slice of order IDs.
- Add new orders dynamically.
- If total orders exceed a constant `MaxOrders`, print a warning.

## 🔹 1. Student Marks – Average

```go
package main
import "fmt"

func main() {
    const Subjects = 5
    marks := []int{80, 75, 90, 85, 70}
    var sum int

    for _, m := range marks {
        sum += m
    }
    avg := float64(sum) / Subjects
    fmt.Println("Average Marks:", avg)
}

```

---

## 🔹 2. Shopping Cart – Discount

```go
package main
import "fmt"

func main() {
    prices := []float64{100, 200, 150}
    var discount float64 = 0.1 // 10%

    var total float64
    for _, price := range prices {
        total += price
    }

    total = total - (total * discount)
    fmt.Println("Total after discount:", total)
}

```

---

## 🔹 3. Employee List – Append New Employee

```go
package main
import "fmt"

func main() {
    employees := []string{"Alice", "Bob"}
    var newEmployee = "Charlie"

    employees = append(employees, newEmployee)
    fmt.Println("Employee List:", employees)
}

```

---

## 🔹 4. Todo Tasks – MaxTasks Limit

```go
package main
import "fmt"

func main() {
    const MaxTasks = 3
    tasks := []string{}
    taskInputs := []string{"Buy milk", "Pay bills", "Call mom", "Extra task"}

    for _, t := range taskInputs {
        if len(tasks) < MaxTasks {
            tasks = append(tasks, t)
        } else {
            fmt.Println("Reached MaxTasks! Cannot add:", t)
        }
    }
    fmt.Println("Tasks:", tasks)
}

```

---

## 🔹 5. Temperature Records – Max & Min

```go
package main
import "fmt"

func main() {
    temps := []int{32, 28, 35, 30, 31}
    max, min := temps[0], temps[0]

    for _, t := range temps {
        if t > max {
            max = t
        }
        if t < min {
            min = t
        }
    }
    fmt.Println("Max Temp:", max, "Min Temp:", min)
}

```

---

## 🔹 6. Bank Transactions – Deduct Fee

```go
package main
import "fmt"

func main() {
    transactions := []float64{500, 1200, 300}
    const Fee = 10.0

    for i, t := range transactions {
        transactions[i] = t - Fee
    }
    fmt.Println("Transactions after fee:", transactions)
}

```

---

## 🔹 7. Top Scorers – Sort & Pick Top 3

```go
package main
import (
    "fmt"
    "sort"
)

func main() {
    scores := []int{55, 90, 85, 72, 60}
    sort.Sort(sort.Reverse(sort.IntSlice(scores)))

    const TopN = 3
    top := scores
    if len(scores) > TopN {
        top = scores[:TopN]
    }
    fmt.Println("Top Scorers:", top)
}

```

---

## 🔹 8. Library Books – Remove by Index

```go
package main
import "fmt"

func main() {
    books := []string{"Book A", "Book B", "Book C"}
    var removeIndex = 1 // remove "Book B"

    books = append(books[:removeIndex], books[removeIndex+1:]...)
    fmt.Println("Books after removal:", books)
}

```

---

## 🔹 9. Student Grades – Pass/Fail

```go
package main
import "fmt"

func main() {
    marks := []int{45, 80, 67, 30, 90}
    const PassingMark = 40

    var passed, failed int
    for _, m := range marks {
        if m >= PassingMark {
            passed++
        } else {
            failed++
        }
    }
    fmt.Println("Passed:", passed, "Failed:", failed)
}

```

---

## 🔹 10. Online Orders – MaxOrders Limit

```go
package main
import "fmt"

func main() {
    const MaxOrders = 5
    orders := []int{}
    newOrders := []int{101, 102, 103, 104, 105, 106}

    for _, o := range newOrders {
        orders = append(orders, o)
        if len(orders) > MaxOrders {
            fmt.Println("⚠️ MaxOrders exceeded! Current orders:", orders)
            break
        }
    }
    fmt.Println("Orders:", orders)
}

```