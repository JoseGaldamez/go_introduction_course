package main

import "fmt"

func main() {
	printExample01()

	rol1 := 1
	rol2 := 1 << 1
	rol3 := 1 << 2
	rol4 := 1 << 3

	fmt.Printf("Roles Binary: %.4b, %.4b, %.4b & %.4b\n", rol1, rol2, rol3, rol4)
	fmt.Printf("Roles Number: %.4d, %.4d, %.4d & %.4d\n", rol1, rol2, rol3, rol4)

	myProfile := rol1 + rol4 + rol2

	fmt.Printf("My profile: (%.4d, %.4b)\n", myProfile, myProfile)

	fmt.Println(0 != (myProfile & rol1))
	fmt.Println(0 != (myProfile & rol2))
	fmt.Println(0 != (myProfile & rol3))
	fmt.Println(0 != (myProfile & rol4))
}

func printExample01() {
	// Operadores binarios
	var a uint16 = 213

	fmt.Printf("Number %d in binary is %b\n", a, a)
	fmt.Printf("\nLeft shift\n")
	fmt.Printf("Number %d in binary is %b\n", a<<1, a<<1)
	fmt.Printf("Number %d in binary is %b\n", a<<10, a<<10)
	fmt.Printf("")
	fmt.Printf("\nRight shift\n")
	fmt.Printf("Number %d in binary is %b\n", a>>1, a>>1)
	fmt.Printf("Number %d in binary is %b\n", a>>10, a>>10)
	fmt.Printf("Number %d in binary is %b\n", a>>5, a>>5)

	// Operadores logicos

	var b uint16 = 20
	fmt.Printf("'a': %.3d - %.10b\n", a, a)
	fmt.Printf("'b': %.3d - %.10b\n\n", b, b)
	fmt.Printf("Bitwise AND: %d - %.10b\n", a&b, a&b)
	fmt.Printf("Bitwise OR: %d - %.10b\n", a|b, a|b)
	fmt.Printf("Bitwise XOR: %d - %.10b\n", a^b, a^b)
	fmt.Printf("Bitwise NOT: %d - %.10b\n", ^a, ^a)
	fmt.Println()
	fmt.Printf("Bitwise NAND: %d - %.10b\n", ^(a & b), ^(a & b))
	fmt.Printf("Bitwise NOR: %d - %.10b\n", ^(a | b), ^(a | b))
	fmt.Printf("Bitwise XNOR: %d - %.10b\n", ^(a ^ b), ^(a ^ b))
}
