package main

import "fmt"

func main() {

	fmt.Println("====== Tarea 1 ====")
	license := false
	age := 18

	switch {
	case license && age >= 15:
		fmt.Println("Puede seguir avanzando")
	default:
		fmt.Println("No puede seguir circulando")
	}

	fmt.Println("====== Tarea 2 ====")
	array := [5]int{4, 2, 5, 6, 7}
	for index, value := range array {
		array[index] = value + 20
	}

	fmt.Println(array)

	fmt.Println("====== Tarea 3 ====")

	var array3 []int
	for {
		var number int
		fmt.Scanln(&number)

		if number == 0 {
			break
		}
		array3 = append(array3, number)
	}

	fmt.Println(array3)

	fmt.Println("====== Tarea 4 ======")

	mapBase := map[string]string{
		"10": "notebook",
		"15": "TV",
		"21": "heladera",
		"27": "monitos",
		"35": "camara",
	}

	var mapResponse = make(map[string]string)
	var arrayResponse []string

	for {
		var codeSelected string
		fmt.Scanln(&codeSelected)
		if codeSelected == "0" {
			break
		}

		product, exist := mapBase[codeSelected]
		if exist {
			mapResponse[codeSelected] = product
			arrayResponse = append(arrayResponse, product)
		} else {
			mapResponse[codeSelected] = "No encontrado"
			arrayResponse = append(arrayResponse, "No encontrado")
		}

	}

	fmt.Println("Los productos seleccionados fueron: ", mapResponse)
	fmt.Println("Los valores del array son: ", arrayResponse)

}
