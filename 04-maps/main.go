package main

import "fmt"

func main() {

	// primer parametro llave, segundo el valor
	map1 := make(map[int]string)
	map1[1] = "A"
	map1[2] = "B"
	map1[3] = "C"
	map1[4] = "D"

	fmt.Println(map1)
	fmt.Println(map1[1])

	// Borrar elementos del mapa
	delete(map1, 2)
	fmt.Println(map1)

	map1[1] = "AA"
	fmt.Println(map1)

	v, exist := map1[7]

	fmt.Println("valor: ", v, " exist: ", exist)

	// Crear mapas sin make
	map2 := map[int]string{
		1: "Hola",
		2: "Mundo",
	}

	fmt.Println(map2)

}
