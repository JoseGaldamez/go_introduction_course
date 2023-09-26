package main

import "fmt"

func main() {

	// IF
	yearsOld := 20

	if yearsOld > 18 {
		fmt.Printf("%d is higher than 18\n", yearsOld)
	}

	// IF ELSE
	boolVar := false

	if boolVar {
		fmt.Println("is true")
	} else {
		fmt.Println("is false")
	}

	// IF := fn()
	if value := true; value {
		fmt.Println("is true")
	}

	// IF anidados
	number := 3
	if number == 1 {
		fmt.Println("One")
	} else if number == 2 {
		fmt.Println("Two")

	} else if number == 3 {
		fmt.Println("Three")
	}

	// SWITCH
	switch number {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Thre")
	default:
		fmt.Println("Undefined")
	}

	// SWITCH without a variable to analysis
	switch {
	case number > 0:
		fmt.Println("es positivo")
	case number < 0:
		fmt.Println("es negativo")
	default:
		fmt.Println("Es igual a cero")
	}

}
