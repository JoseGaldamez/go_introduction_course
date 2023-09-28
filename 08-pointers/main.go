package main

import "fmt"

func main() {
	var i int
	var ipointer *int

	i = 34
	ipointer = &i

	fmt.Println(i)
	fmt.Println(&i)

	fmt.Println(ipointer)
	fmt.Println(*ipointer)

	fmt.Println()
	// Moficicar un puntero modifica la variable a la que se esta refenciando
	*ipointer = 1

	fmt.Println(i)
	fmt.Println(&i)

	fmt.Println(ipointer)
	fmt.Println(*ipointer)

	fmt.Println()

	myVar := 30
	fmt.Printf("Address: %p, values: %v\n", &myVar, myVar)

	increment(myVar)
	fmt.Printf("Address: %p, values: %v\n", &myVar, myVar)

	incrementPointer(&myVar)
	fmt.Printf("Address: %p, values: %v\n", &myVar, myVar)

	fmt.Println()

	var mySlice []int
	mySlice = append(mySlice, 3, 4, 2)

	fmt.Printf("Address 1: %p, Address 2: %p, Address 3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])
	fmt.Println(mySlice)
	updateSlice(mySlice)
	fmt.Println()
	fmt.Println(mySlice)
	fmt.Printf("Address 1: %p, Address 2: %p, Address 3: %p\n", &mySlice[0], &mySlice[1], &mySlice[2])

	fmt.Println()

	myStruc := MyStruct{ID: 1, Name: "Jose"}
	fmt.Printf("address: %p\n", &myStruc)
	fmt.Printf("id: %d, name: %s\n", myStruc.ID, myStruc.Name)
	fmt.Println()

	updateStruct(&myStruc)
	fmt.Printf("id: %d, name: %s\n", myStruc.ID, myStruc.Name)

	fmt.Println()

	mysecondStruc := MySecondStruc{ID: 1, Name: "Jose"}

	fmt.Printf("Address of struct: %p\n", &mysecondStruc)
	fmt.Println(mysecondStruc)
	mysecondStruc.SetMethod("No pointed")
	fmt.Println(mysecondStruc)
	mysecondStruc.SetMethodPointed("Pointed")
	fmt.Println(mysecondStruc)

}

func increment(value int) {
	fmt.Println(&value)
	value++
}

func incrementPointer(value *int) {
	fmt.Println(value)
	*value++
}

func updateSlice(values []int) {
	fmt.Printf("Address 1: %p, Address 2: %p, Address 3: %p\n", &values[0], &values[1], &values[2])

	values[0] = 12
}

type MyStruct struct {
	ID   uint
	Name string
}

func updateStruct(v *MyStruct) {
	fmt.Printf("address in function: %p\n", v)
	v.ID = 999
	v.Name = "update struct"
}

type MySecondStruc struct {
	ID   uint
	Name string
}

func (mysecondStruct MySecondStruc) SetMethod(value string) {
	fmt.Println(value)
	fmt.Printf("Address: %p\n", &mysecondStruct)
	mysecondStruct.Name = value
}

func (mysecondStruct *MySecondStruc) SetMethodPointed(value string) {
	fmt.Println(value)
	fmt.Printf("Address: %p\n", mysecondStruct)
	mysecondStruct.Name = value
}
