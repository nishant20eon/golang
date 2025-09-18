# Arrays + Variables + Constants

## 🔹 10 Practice Questions: Arrays + Variables + Constants

1. **Declare a constant for array size. Create an integer array of that size and print all elements (default values).**
2. **Initialize an array with constant values, then use a variable loop counter to print all elements.**
3. **Write a program to calculate the sum of elements in an integer array. Use a constant for array length.**
4. **Declare a string array of constant size. Use a variable to assign values at runtime and print the array.**
5. **Create a constant for the number of subjects, declare an array for marks, and calculate the average using a variable sum.**
6. **Use a constant to define array size, fill it with squares of indices using a loop with a variable index, and print it.**
7. **Demonstrate shadowing: declare a global constant for array size, then inside `main()` redeclare a variable with the same name and use it to fill the array.**
8. **Use a constant array size, declare an array, and reverse it using variable swapping.**
9. **Store a set of constant values (e.g., weekdays) in an array and use a variable index to access and print one element.**
10. **Declare a constant for the maximum array size, but initialize only part of it using a variable. Print both initialized and default elements.**

## 🔹 1. Constant array size + default values

```go
package main
import "fmt"

func main() {
    const Size = 5
    var arr [Size]int // default values = 0
    fmt.Println("Array with default values:", arr)
}

```

---

## 🔹 2. Initialize array with constants + print using variable loop

```go
package main
import "fmt"

func main() {
    arr := [3]int{10, 20, 30}

    for i := 0; i < len(arr); i++ { // variable loop counter
        fmt.Println("Element", i, ":", arr[i])
    }
}

```

---

## 🔹 3. Sum of elements using constant length

```go
package main
import "fmt"

func main() {
    const Size = 4
    arr := [Size]int{5, 10, 15, 20}
    var sum int

    for i := 0; i < Size; i++ {
        sum += arr[i]
    }
    fmt.Println("Sum:", sum)
}

```

---

## 🔹 4. String array + assign using variable

```go
package main
import "fmt"

func main() {
    const Size = 3
    var names [Size]string

    var input string
    for i := 0; i < Size; i++ {
        input = fmt.Sprintf("Name%d", i+1) // variable assignment
        names[i] = input
    }

    fmt.Println("Names array:", names)
}

```

---

## 🔹 5. Average marks using const subjects

```go
package main
import "fmt"

func main() {
    const Subjects = 5
    marks := [Subjects]int{80, 75, 90, 85, 70}
    var sum int

    for i := 0; i < Subjects; i++ {
        sum += marks[i]
    }

    avg := float64(sum) / Subjects
    fmt.Println("Average Marks:", avg)
}

```

---

## 🔹 6. Squares of indices

```go
package main
import "fmt"

func main() {
    const Size = 6
    var arr [Size]int

    for i := 0; i < Size; i++ {
        arr[i] = i * i
    }

    fmt.Println("Squares array:", arr)
}

```

---

## 🔹 7. Shadowing constant with variable

```go
package main
import "fmt"

const Size = 4

func main() {
    fmt.Println("Global constant Size:", Size)

    Size := 3 // shadowing with variable
    arr := make([]int, Size)

    for i := 0; i < Size; i++ {
        arr[i] = i + 1
    }

    fmt.Println("Local variable Size:", Size)
    fmt.Println("Array:", arr)
}

```

---

## 🔹 8. Reverse array using variable swapping

```go
package main
import "fmt"

func main() {
    const Size = 5
    arr := [Size]int{1, 2, 3, 4, 5}

    for i := 0; i < Size/2; i++ {
        arr[i], arr[Size-1-i] = arr[Size-1-i], arr[i]
    }

    fmt.Println("Reversed array:", arr)
}

```

---

## 🔹 9. Weekdays array + variable access

```go
package main
import "fmt"

func main() {
    weekdays := [7]string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

    var index = 3
    fmt.Println("Day at index", index, ":", weekdays[index])
}

```

---

## 🔹 10. Partially filled array with const max size

```go
package main
import "fmt"

func main() {
    const MaxSize = 6
    var arr [MaxSize]int

    usedSize := 3 // variable
    for i := 0; i < usedSize; i++ {
        arr[i] = (i + 1) * 10
    }

    fmt.Println("Partially filled array:", arr)
}

```