package main

import "fmt"

func main() {

	yearsOld := 32

	fmt.Println("Operators")
	fmt.Println(yearsOld > 30)  // true
	fmt.Println(yearsOld < 30)  // false
	fmt.Println(yearsOld <= 32) // true
	fmt.Println(yearsOld >= 32) // true
	fmt.Println(yearsOld == 32) // true

	fmt.Println()
	// OR
	fmt.Println(yearsOld < 32 || yearsOld == 32) // (false, true) true
	fmt.Println(yearsOld < 32 || yearsOld == 33) // (true, false) true
	fmt.Println(yearsOld > 32 || yearsOld == 33) // (false, false) false

	fmt.Println()
	// AND
	fmt.Println(yearsOld < 32 && yearsOld == 32) // (false, true) false
	fmt.Println(yearsOld < 32 && yearsOld == 33) // (true, false) false
	fmt.Println(yearsOld > 32 && yearsOld == 33) // (false, false) false

	fmt.Println()
	// NOT !
	fmt.Println(!true)            // false
	fmt.Println(!false)           // true
	fmt.Println(!(yearsOld > 30)) // false

	fmt.Println()
	fmt.Println(yearsOld < 25 && yearsOld == 32 || yearsOld < 40)   // (false and true of true) true
	fmt.Println(yearsOld < 25 && (yearsOld == 32 || yearsOld < 40)) // (false and true) false

}
