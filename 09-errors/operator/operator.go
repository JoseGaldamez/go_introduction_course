package operator

import "fmt"

func Div(x, y float64) (z float64) {

	defer func() {
		fmt.Println("In the defert function")
		r := recover()

		if r != nil {
			fmt.Println("An error happend")
		}
	}()

	z = x / y

	return z

}
