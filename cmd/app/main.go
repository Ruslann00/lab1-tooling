// Package main is the entry point of the application.
package main

import (
	"fmt"
	"log"

	"github.com/Ruslann00/lab1-tooling/internal"
)

func main() {
	sum := internal.Add(10, 5)
	difference := internal.Subtract(10, 5)

	result, err := internal.Divide(10, 2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Division:", result)
}
