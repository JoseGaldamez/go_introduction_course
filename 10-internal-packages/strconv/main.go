package main

import (
	"fmt"
	"strconv"
)

func main() {
	st := strconv.Itoa(-43)
	fmt.Println(st)

	// Convertir flotante enviando el tipo de datos que necesitamos
	st = strconv.FormatFloat(3.1426, 'E', -1, 64)
	fmt.Println(st)

	// Convertir entero enviando el tipo de datos que necesitamos
	st = strconv.FormatInt(-42, 8)
	fmt.Println(st)

	// Convertir entero positivo enviando el tipo de datos que necesitamos
	st = strconv.FormatUint(42, 10)
	fmt.Println(st)

	// Parseos
	b, err := strconv.ParseBool("true")
	fmt.Println(b, err)

	f, err := strconv.ParseFloat("3.1416", 64)
	fmt.Println(f, err)

	in, err := strconv.ParseInt("42", 10, 64)
	fmt.Println(in, err)

	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(u, err)

}
