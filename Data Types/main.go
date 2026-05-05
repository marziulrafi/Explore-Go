package main

import "fmt"

func main() {

	// Integer
	var a int = 42       // default int (platform-dependent: 32 or 64 bit)
	var b int8 = -128    // 8-bit signed integer (-128 to 127)
	var c uint16 = 65535 // 16-bit unsigned integer (0 to 65535)
	fmt.Println("Integers:", a, b, c)

	// Floating-point
	var f1 float32 = 3.14
	var f2 float64 = 2.718281828459045
	fmt.Println("Floats:", f1, f2)

	// Complex numbers
	var complexNum complex64 = 1 + 2i
	fmt.Println("Complex:", complexNum)
	fmt.Println("Real part:", real(complexNum))
	fmt.Println("Imag part:", imag(complexNum))


	// Boolean
	var isTrue bool = true
	var isFalse bool = false
	fmt.Println("Booleans:", isTrue, isFalse)

	// Boolean expressions
	fmt.Println("Is 10 > 5?", 10 > 5)
	fmt.Println("Is 3 == 7?", 3 == 7)


	// String
	var name string = "Marziul"
	message := "Learning Go is fun!"
	fmt.Println("Strings:", name, "-", message)

	// String concatenation
	fullMsg := name + " says: " + message
	fmt.Println("Concatenated:", fullMsg)

	// Multi-line string
	multiLine := `
This is a 
multi-line 
string in Go.`
	fmt.Println("Multiline:", multiLine)


	// Arrays
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	// Access and modify
	arr[1] = 5
	fmt.Println("Modified array:", arr)
	fmt.Println("Array length:", len(arr))

	
}
