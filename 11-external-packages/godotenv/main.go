package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	fmt.Println(os.Getenv("MY_ENV_V1"))
	fmt.Println(os.Getenv("MY_ENV_V2"))
	fmt.Println(os.Getenv("MY_ENV_V3"))
	fmt.Println(os.Getenv("MY_ENV_V4"))

	if err := godotenv.Load("otherFolder/.var"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(os.Getenv("MY_ENV_V1"))
	fmt.Println(os.Getenv("MY_ENV_V2"))
	fmt.Println(os.Getenv("MY_ENV_V3"))
	fmt.Println(os.Getenv("MY_ENV_V4"))

	myEnv, err := godotenv.Read("otherFolder/.var")
	fmt.Println(myEnv)
	fmt.Println(err)

	if err := godotenv.Overload(); err != nil {
		fmt.Println()
	}

	fmt.Println(os.Getenv("MY_ENV_V1"))
	fmt.Println(os.Getenv("MY_ENV_V2"))
	fmt.Println(os.Getenv("MY_ENV_V3"))
	fmt.Println(os.Getenv("MY_ENV_V4"))

}
