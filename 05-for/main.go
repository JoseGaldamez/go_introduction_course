package main

import "fmt"

func main() {

	sum := 0

	// for completo
	for i := 0; i < 10; i++ {
		sum++
		fmt.Println("Valor de sum: ", sum)
	}

	// for solo con condicional
	for sum < 20 {
		sum++
	}

	fmt.Println(sum)

	// for infinito con interrupción
	for {
		if sum >= 100 {
			break
		}

		sum++
	}

	fmt.Println(sum)

	// recorrer Arrays
	arr := []int{1, 2, 3, 4, 5}

	for index, value := range arr {
		fmt.Println("Index: ", index, " valor: ", value)
	}

	for _, value := range arr {
		fmt.Println("Index no interesa: _ valor: ", value)
	}

	// recorrer Arrays
	map1 := map[string]string{"HN": "Honduras", "GT": "Guatemala"}

	for key, value := range map1 {
		fmt.Println("Key: ", key, " valor: ", value)
	}

	map2 := map[string][]int{
		"A": nil,
		"B": {1, 2, 3, 4},
		"C": {5, 6, 7, 8},
	}

	for key, value := range map2 {
		fmt.Println("Llave del map: ", key)
		for _, v := range value {
			fmt.Println("Valores del array", v)
		}
		fmt.Println()
	}

}
