package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func main() {
	v1 := []float64{1.3, 5.4}
	v2 := []int32{9, 23, 5, 6}

	fmt.Println(Sum01(v1))
	fmt.Println(Sum01(v2))
	fmt.Println()

	fmt.Println(Sum02(v1))
	fmt.Println(Sum02(v2))

	fmt.Println()

	anyType("Hola mundo", "Otra cosa")
	anyType(2, 3)
	//anyType(2,"Texto") // el generico se setea como el primer tipo de datos detectado
	anyTypeTwo(2, "Aqui si se puede")

	fmt.Println()

	comparablesAnyType(2, 4)
	comparablesAnyType(4, 4)

	orderedValues(2, 4)

	csInt := CustomSlice[int]{1, 2, 3, 4}
	fmt.Println(csInt)

	csString := CustomSlice[string]{"Algo", "interesante", "de", "Golang"}
	fmt.Println(csString)

	x, y := Point(5), Point(2)

	fmt.Println(MinNumber(x, y))

	fg := MyGenericStruc[MyFirstData]{StringValue: "Test", Data: MyFirstData{}}
	fg.Data.PrintOne()

	sg := MyGenericStruc[MySecondData]{StringValue: "Test", Data: MySecondData{}}
	sg.Data.PrintTwo()

}

func Sum01[N int | int32 | int64 | float64 | float32](v []N) N {
	var total N

	for _, tV := range v {
		total += tV
	}

	return total
}

type Number interface {
	int | int32 | int64 | float64 | float32
}

func Sum02[N Number](v []N) N {
	var total N

	for _, tV := range v {
		total += tV
	}

	return total
}

func anyType[N any](v1, v2 N) {
	fmt.Println(v1, v2)
	// if v1 == v2 { // valores de tipo any no son comparables
	// 	fmt.Println("Son iguales")
	// }
}

func anyTypeTwo[N any, V any](v1 N, v2 V) {
	fmt.Println(v1, v2)
}

func comparablesAnyType[N comparable](v1, v2 N) {
	fmt.Println(v1, v2)
	if v1 == v2 { // valores de tipo comparable si se pueden comparar, solo == o !=
		fmt.Println("Son iguales")
	} else {
		fmt.Println("No son iguales")
	}
}

func orderedValues[N constraints.Ordered](v1, v2 N) {
	fmt.Println(v1, v2)
	fmt.Println(v1 == v2)
	fmt.Println(v1 < v2)
	fmt.Println(v1 > v2)
}

type CustomSlice[V int | string] []V

func MinNumber[T NumberCompatible](x, y T) T {
	if x < y {
		return x
	}

	return y
}

type Point int

type NumberCompatible interface {
	~int | ~int32 | ~int64 | ~float64 | ~float32 // permite tipos de datos o interfaces que implementan los tipos
}

type (
	MyGenericStruc[D any] struct {
		StringValue string
		Data        D
	}

	MyFirstData struct {
	}

	MySecondData struct{}
)

func (d MyFirstData) PrintOne() {
	fmt.Println("Print first")
}

func (d MySecondData) PrintTwo() {
	fmt.Println("Print second")
}
