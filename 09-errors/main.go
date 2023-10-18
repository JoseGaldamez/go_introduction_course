package main

import (
	"errors"
	"fmt"

	"github.com/JoseGaldamez/go_introduction_course/09-errors/operator"
)

func main() {
	var err error
	fmt.Println(err) // Valor por defecto es nil

	err = errors.New("My new error")
	fmt.Println(err)

	fmt.Println(err.Error()) // regresa el string del error

	err2 := fmt.Errorf("El valor %s is not valid as a number: %d.", "mystring", 1)
	fmt.Println(err2)
	fmt.Printf("%T\n", err2)
	fmt.Println(err2.Error())

	defer func() {
		// a function with defer is called at the end of a function
		fmt.Println("In my defert")

		// recover returns a error when the function end with a error
		r := recover()
		if r != nil {
			fmt.Println("Recovered in: ", r)
		}
	}()

	var x float64 = 3
	var y float64 = 2
	z := operator.Div(x, y)

	fmt.Println("Result: ", z)

}
