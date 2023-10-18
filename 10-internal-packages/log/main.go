package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	log.Println("Hola mundo")

	// log.Fatal("La aplicación se ha roto.")
	// log.Panic("La aplicación ha entrado en pánico")

	filename := fmt.Sprintf("%d.log", time.Now().UnixNano())
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	logObject := log.New(file, "Logger file: ", log.LstdFlags|log.Lshortfile)

	logObject.Println("test 01")
	logObject.Println("test 02")
	logObject.Println("test 03")

	// cambiar el prefix en tiempo de ejecución
	logObject.SetPrefix("error: ")

	logObject.Println("test 04")
	logObject.Println("test 05")
	logObject.Println("test 06")

}
