package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {

	uid1 := uuid.New().String()
	fmt.Println(uid1)

	uid2 := uuid.NewString()
	fmt.Println(uid2)

	uid3 := uuid.New()
	fmt.Println(uid3.Version())

}
