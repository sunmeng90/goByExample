package main

import (
	"fmt"
)

func init() {
	println("init")
}

func main() {
	println("hello world !")

	var firstName string = "Meng"
	var lastName string = "Sun"

	fmt.Printf("Hello, %s, %s\n", firstName, lastName)

	var firstNamePointer *string = &firstName
	fmt.Printf("Pointer: %T\n", firstNamePointer)
	fmt.Printf("Mem address: %p\n", firstNamePointer)

	var pValue string = *firstNamePointer

	fmt.Printf("Pointer value type: %T\n", pValue)
	fmt.Printf("Pointer value: %s\n", pValue)

}
