package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("myFile.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data := make([]byte, 100)
	c, err := file.Read(data)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("read: %d bytes: %q\n", c, data[:c])

	fmt.Println("========= variables de entorno =========")

	// leer variables de entorno
	env1 := os.Getenv("Path")
	fmt.Println(env1)

	// setear una variable de entorno
	os.Setenv("MI_ENV", "MI_VALOR")
	fmt.Println(os.Getenv("MI_ENV"))

	dir, _ := os.Getwd()
	fmt.Println(dir)

	fmt.Println(os.Getenv("ENV_NOT_EXIST"))

	environ, exist := os.LookupEnv("ENV_NOT_EXIST")
	fmt.Println(environ, exist)

	fmt.Println()

	os.Setenv("DB_USERNAME", "jose")
	os.Setenv("DB_PASS", "jose")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "27018")
	os.Setenv("DB_NAME", "jose")

	// nos permite usar ${} o simplemente $ para acceder a variables de entorno
	dbUrl := os.ExpandEnv("mongo://${DB_USERNAME}:${DB_PASS}@$DB_HOST:$DB_PORT/$DB_NAME")
	fmt.Println(dbUrl)

}
