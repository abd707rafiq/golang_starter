package main // package declaration

import (
	"fmt" /// fmt package
)

// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

/// helo world in golang

// func main() { // go function
// 	fmt.Println("Hello World!") /// print
// }

//// go variables

// In Go, there are different types of variables, for example:

// int- stores integers (whole numbers), such as 123 or -123
// float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
// string - stores text, such as "Hello World". String values are surrounded by double quotes
// bool- stores values with two states: true or false

/// declaring go variables

// 1- with var ketword

// var variableName type=value
// Always have to specify type or value or both

// 2- with := sign

// variablename := value

// In this case, the type of the variable is inferred from the value (means that the compiler decides the type of the variable, based on the value).

// Variable Declaration With Initial Value
// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<
// func main() {

// 	var student1 string = "john"
// 	var student2 = "ali"
// 	x := 2

// 	fmt.Println(student1)
// 	fmt.Println(student2)
// 	fmt.Println(x)

// }
// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

// Variable Declaration Without Initial Value

// func main() {
// 	var a string
// 	var b int
// 	var c bool

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }

// Value Assignment After Declaration
// <<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<

// func main() {
// 	var student1 string
// 	student1 = "John"
// 	fmt.Println(student1)
// }

// Note: It is not possible to declare a variable using ":=" without assigning a value to it.

// Go Multiple Variable Declaration

// go allows to declare multiple variblaes

func main() {
	var a, b, c, d int = 1, 3, 5, 7

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// Note: If you use the type keyword, it is only possible to declare one type of variable per line.
// If the type keyword is not specified, you can declare different types of variables on the same line:
