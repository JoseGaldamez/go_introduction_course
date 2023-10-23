package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID        int64    `myLabel:"lb1" myOtherLabel:"lob2"`
	Email     string   `myLabel:"lb2"`
	FirrtName string   `myLabel:"lb3"`
	LastName  string   `myLabel:"lb4"`
	Age       *float64 `myLabel:"lb5"`
	Address   Address  `myLabel:"lb6"`
}

type Address struct {
	Country string
	City    string
}

func main() {
	myInt := 5
	Examine(myInt)

	myPtr := &myInt
	Examine(myPtr)

	var age float64 = 32
	user := User{1, "josegaldamez@gmail.com", "Jose", "Galdamez", &age, Address{"Honduras", "El Progreso"}}
	Examine(user)

}

func Examine(data interface{}) {

	if reflect.ValueOf(data).Kind() == reflect.Struct {
		v := reflect.ValueOf(data)
		t := reflect.TypeOf(data)
		fmt.Println("Number of fields; ", v.NumField())

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field: %d, type: %s, value: %v\n", i, v.Field(i).Kind(), v.Field(i))
			c := t.Field(i).Tag
			fmt.Println("Tag: ", c.Get("myLabel"))
			fmt.Println("Tag: ", c.Get("myOtherLabel"))
		}

		return
	}

	t := reflect.TypeOf(data)
	v := reflect.ValueOf(data)
	k := t.Kind()

	fmt.Println("=================")
	fmt.Println("Type: ", t)
	fmt.Println("Value: ", v)
	fmt.Println("Kind: ", k)
	fmt.Println()
}
