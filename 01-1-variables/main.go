package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

func main() {

	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUnintVar uint // is equal a int, but must be a positive value
	myUnintVar = 12
	fmt.Println("My unint variable is: ", myUnintVar)

	var myStringVar string
	myStringVar = "My string variable"
	fmt.Println(myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println("My bool variable: ", myBoolVar)

	fmt.Println("My variable address is: ", &myStringVar)

	myIntVar2 := 12
	fmt.Println("My variable is: ", myIntVar2)

	fmt.Println()

	const myIntConst int = 12
	fmt.Println("My const is", myIntConst)

	const myStringConst string = "Const variable"
	// myStringConst = "Other value" // this is not admited
	fmt.Println("My const string is", myStringConst)

	fmt.Println()

	var my8bitsVar int8
	fmt.Printf("int8 default value %d\n", my8bitsVar)
	fmt.Printf("int8 default value %s\n", "my8bitsVar")
	fmt.Printf("int8 default value %f\n", 0.5)

	fmt.Println()

	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my8bitsVar, unsafe.Sizeof(my8bitsVar), unsafe.Sizeof(my8bitsVar)*8)

	var my16bitsvar int16
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my16bitsvar, unsafe.Sizeof(my16bitsvar), unsafe.Sizeof(my16bitsvar)*8)

	var my32bitsvar int32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my32bitsvar, unsafe.Sizeof(my32bitsvar), unsafe.Sizeof(my32bitsvar)*8)

	var my64bitsvar int64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", my64bitsvar, unsafe.Sizeof(my64bitsvar), unsafe.Sizeof(my64bitsvar)*8)

	var myuintBitsvar uint
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myuintBitsvar, unsafe.Sizeof(myuintBitsvar), unsafe.Sizeof(myuintBitsvar)*8)

	fmt.Println()

	var myFloat32Var float32
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myFloat64Var float64
	fmt.Printf("type: %T, bytes: %d, bits: %d\n", myFloat64Var, unsafe.Sizeof(myFloat64Var), unsafe.Sizeof(myFloat64Var)*8)

	fmt.Println()
	var myStringVar3 string
	fmt.Printf("String default value: %s\n", myStringVar3)

	myStringVar5 := `My string variable in golang
	with multiple
	line`

	fmt.Printf("%s\n", myStringVar5)

	// Esto genera un scope, básicamente es para no tener
	// conflicto con los nombres de variables generados anteriormente
	{
		fmt.Println()
		floarVar := 33.11
		fmt.Printf("type: %T, value: %.2f\n", floarVar, floarVar)
		floatStringVar := fmt.Sprintf("%.2f", floarVar)
		fmt.Printf("type: %T, value: %s\n", floatStringVar, floatStringVar)

		fmt.Println()
		intVar := 33
		fmt.Printf("type: %T, value: %.2d\n", intVar, intVar)
		intStringVar := fmt.Sprintf("%d", intVar)
		fmt.Printf("type: %T, value: %s\n", intStringVar, intStringVar)

		fmt.Println()
		intValue, err := strconv.ParseInt("100", 0, 64)
		fmt.Println(err)
		fmt.Printf("type: %T, Error: %d", intValue, intValue)

		fmt.Println()
		intValue2, err2 := strconv.ParseInt("Hola", 0, 64)
		fmt.Println(err2)
		fmt.Printf("type: %T, Error: %d\n", intValue2, intValue2)

		fmt.Println()
		floatVar1, _ := strconv.ParseFloat("-11.02", 64)
		fmt.Printf("type: %T, Error: %f", floatVar1, floatVar1)

	}

}
