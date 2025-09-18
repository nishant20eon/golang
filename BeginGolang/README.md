
# BeginGolang

A simple Go project to help you get started with Go basics, including project setup and variable usage.

---

## Project Setup

1. **Initialize the Go module:**
    ```sh
    go mod init begingolang
    ```

2. **Run the project:**
    ```sh
    go run .
    # or
    go run ./main.go
    ```

---

## Go Variable Basics

Variables are used to store data in Go. You must declare the type when using the `var` keyword.

**Syntax:**
```go
var varName varType = value
```

**Example:**
```go
var agency string = "Fast Track"
```

**Short declaration:**
```go
name := "nishant"
```

### Common Data Types
| Type     | Declaration Example                | Default Value |
|----------|------------------------------------|--------------|
| string   | var s string = "hello"            | ""           |
| int      | var i int = 10                    | 0            |
| float64  | var f float64 = 3.14              | 0.0          |
| bool     | var b bool = true                 | false        |

---

## Example Output


```
Hello, World!
Welcome to Fast Track
My name is nishant
My full name is nishant kumar
Total cars available are 10
Starting price of car is 20000
Is car available? true
length of empName is 8
Hello, World!
false
1
ll
[104 101 108 108 111]

--- Strings Functions Examples ---
Original: '  Hello, GoLang World!  '
TrimSpace: 'Hello, GoLang World!'
ToLower: hello, golang world!
ToUpper: HELLO, GOLANG WORLD!
Contains 'GoLang': true
ReplaceAll: Hello, Gopher World!
Split by ', ': [Hello GoLang World!]
Join with ' | ': Hello | GoLang World!
HasPrefix 'Hello': true
HasSuffix 'World!': true
Index of 'GoLang': 7
Count of 'l': 3
Repeat 'Go! ': Go! Go! Go! 

--- Numeric Functions Examples ---
String to int: 12345
Kelvin to Fahrenheit: 59.00°F

--- Math Functions Examples ---
Abs(-10): 10
Pow(2, 3): 8
Sqrt(16): 4
Max(10, 20): 20
Min(10, 20): 10
Round(4.7): 5
Ceil(4.3): 5
Floor(4.7): 4
String Builder: Hello, World!
String Builder Length: 13
String Builder Capacity: 32
```

---


---

## Most Used Strings Functions in Go

Here are some of the most commonly used functions from the `strings` package in Go, with examples:

```go
import "strings"

sample := "  Hello, GoLang World!  "
trimmed := strings.TrimSpace(sample)                // "Hello, GoLang World!"
lower := strings.ToLower(trimmed)                   // "hello, golang world!"
upper := strings.ToUpper(trimmed)                   // "HELLO, GOLANG WORLD!"
contains := strings.Contains(trimmed, "GoLang")     // true
replaced := strings.ReplaceAll(trimmed, "GoLang", "Gopher") // "Hello, Gopher World!"
words := strings.Split(trimmed, ", ")              // ["Hello" "GoLang World!"]
joined := strings.Join(words, " | ")               // "Hello | GoLang World!"
hasPrefix := strings.HasPrefix(trimmed, "Hello")    // true
hasSuffix := strings.HasSuffix(trimmed, "World!")   // true
index := strings.Index(trimmed, "GoLang")           // 7
count := strings.Count(trimmed, "l")                // 3
repeat := strings.Repeat("Go! ", 3)                 // "Go! Go! Go! "
```

**Output Example:**
```
Original: '  Hello, GoLang World!  '
TrimSpace: 'Hello, GoLang World!'
ToLower: hello, golang world!
ToUpper: HELLO, GOLANG WORLD!
Contains 'GoLang': true
ReplaceAll: Hello, Gopher World!
Split by ', ': [Hello GoLang World!]
Join with ' | ': Hello | GoLang World!
HasPrefix 'Hello': true
HasSuffix 'World!': true
Index of 'GoLang': 7
Count of 'l': 3
Repeat 'Go! ': Go! Go! Go! 
```

---

Happy Learning Go!


