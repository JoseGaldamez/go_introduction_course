package main

import "fmt"

func main() {

	license := false
	age := 18

	switch {
	case license && age >= 15:
		fmt.Println("Puede seguir avanzando")
	default:
		fmt.Println("No puede seguir circulando")
	}
}
