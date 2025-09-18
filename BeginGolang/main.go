package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Hello, World!\n")

	var agency string = "Fast Track"

	fmt.Printf("Welcome to %s\n", agency)

	var name string = "nishant"
	fmt.Printf("My name is %s\n", name)

	secondName := "nishant kumar"

	fmt.Printf("My full name is %s\n", secondName)

	var totalCars int = 10
	fmt.Printf("Total cars available are %d\n", totalCars)

	startingPrice := 20000
	fmt.Printf("Starting price of car is %d\n", startingPrice)

	var isAvailable bool = true
	fmt.Printf("Is car available? %t\n", isAvailable)

	var empName = "John Doe"

	length := len(empName)

	fmt.Printf("length of empName is %d\n", length)

	str1 := "Hello, "
	str2 := "World!"
	fullStr := str1 + str2
	fmt.Println(fullStr)

	str3 := "hello"

	flag := strings.EqualFold(str3, str1)

	fmt.Println(flag)

	index := strings.Index(str3, "e")
	fmt.Println(index)

	// string slicing
	substr := str3[2:4]
	fmt.Println(substr)

	// string to rune conversion
	runes := []rune(str3)
	fmt.Println(runes)

	// --- Most Used Strings Functions Examples ---
	fmt.Println("\n--- Strings Functions Examples ---")

	sample := "  Hello, GoLang World!  "
	fmt.Printf("Original: '%s'\n", sample)

	// TrimSpace
	trimmed := strings.TrimSpace(sample)
	fmt.Printf("TrimSpace: '%s'\n", trimmed)

	// ToLower and ToUpper
	fmt.Printf("ToLower: %s\n", strings.ToLower(trimmed))
	fmt.Printf("ToUpper: %s\n", strings.ToUpper(trimmed))

	// Contains
	fmt.Printf("Contains 'GoLang': %t\n", strings.Contains(trimmed, "GoLang"))

	// ReplaceAll
	replaced := strings.ReplaceAll(trimmed, "GoLang", "Gopher")
	fmt.Printf("ReplaceAll: %s\n", replaced)

	// Split
	words := strings.Split(trimmed, ", ")
	fmt.Printf("Split by ', ': %v\n", words)

	// Join
	joined := strings.Join(words, " | ")
	fmt.Printf("Join with ' | ': %s\n", joined)

	// HasPrefix and HasSuffix
	fmt.Printf("HasPrefix 'Hello': %t\n", strings.HasPrefix(trimmed, "Hello"))
	fmt.Printf("HasSuffix 'World!': %t\n", strings.HasSuffix(trimmed, "World!"))

	// Index
	fmt.Printf("Index of 'GoLang': %d\n", strings.Index(trimmed, "GoLang"))

	// Count
	fmt.Printf("Count of 'l': %d\n", strings.Count(trimmed, "l"))

	// Repeat
	fmt.Printf("Repeat 'Go! ': %s\n", strings.Repeat("Go! ", 3))

	// numeric

	fmt.Println("\n--- Numeric Functions Examples ---")
	numStr := "12345"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("String to int: %d\n", num)
	}

	// convert 288.15 Kelvin to Fahrenheit
	kelvinStr := "288.15"
	kelvinStr = strings.TrimSpace(kelvinStr)
	kelvin, err := strconv.ParseFloat(kelvinStr, 64)

	if err != nil {
		fmt.Printf("Error converting string to float: %v\n", err)
	} else {
		fahrenheit := (kelvin-273.15)*9/5 + 32
		fmt.Printf("Kelvin to Fahrenheit: %.2f°F\n", fahrenheit)
	}

	// --- Most Used Math Functions Examples ---
	fmt.Println("\n--- Math Functions Examples ---")
	fmt.Printf("Abs(-10): %.0f\n", math.Abs(-10))
	fmt.Printf("Pow(2, 3): %.0f\n", math.Pow(2, 3))
	fmt.Printf("Sqrt(16): %.0f\n", math.Sqrt(16))
	fmt.Printf("Max(10, 20): %.0f\n", math.Max(10, 20))
	fmt.Printf("Min(10, 20): %.0f\n", math.Min(10, 20))
	fmt.Printf("Round(4.7): %.0f\n", math.Round(4.7))
	fmt.Printf("Ceil(4.3): %.0f\n", math.Ceil(4.3))
	fmt.Printf("Floor(4.7): %.0f\n", math.Floor(4.7))

	// String Builder
	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("World!")
	fmt.Println("String Builder:", sb.String())

	fmt.Printf("String Builder Length: %d\n", sb.Len())
	fmt.Printf("String Builder Capacity: %d\n", sb.Cap())

	// write code for type conversion from string to int and float and vice versa
	fmt.Println("\n--- Type Conversion Examples ---")
	intVal := 42
	floatVal := 3.14159
	strFromInt := strconv.I toa(intVal)
	strFromFloat := strconv.FormatFloat(floatVal, 'f', 2, 64)
	fmt.Printf("Integer to String: %s\n", strFromInt)
	fmt.Printf("Float to String: %s\n", strFromFloat)

	// String to Integer
	strToInt, err := strconv.Atoi(strFromInt)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	} else {
		fmt.Printf("String to Integer: %d\n", strToInt)
	}
	// String to Float
	strToFloat, err := strconv.ParseFloat(strFromFloat, 64)
	if err != nil {
		fmt.Println("Error converting string to float:", err)
	} else {
		fmt.Printf("String to Float: %.2f\n", strToFloat)
	}

	str := string(80)

	fmt.Println(str)

	// conatants
	const pi = 3.14
	const (
		statusOK       = 200
		statusNotFound = 404
	)
	fmt.Printf("Constant Pi: %.2f\n", pi)
	fmt.Printf("Status OK: %d, Status Not Found: %d\n", statusOK, statusNotFound)	
	//iota example
	const (
		first  = iota // 0
		second        // 1
		third         // 2
	)
	fmt.Printf("Iota Values: first=%d, second=%d, third=%d\n", first, second, third)
}
