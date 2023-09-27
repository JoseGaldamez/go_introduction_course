package main

import (
	"encoding/json"
	"fmt"

	"github.com/JoseGaldamez/go_introduction_course/07-structs-interfaces/structs"
)

func main() {

	var product1 structs.Product
	product1.ID = 12
	product1.Name = "test"
	fmt.Println(product1)

	product2 := structs.Product{}
	fmt.Println(product2)

	// Esta forma no es recomendada ya que requiere enviar a ciegas todos los valores
	product3 := structs.Product{2, "Test", structs.Type{}, 1, 12.21, nil}
	fmt.Println(product3)

	// Esta forma es un poco más recomenda
	product4 := structs.Product{
		ID:    1,
		Name:  "Test",
		Price: 21.12,
		Type:  structs.Type{},
		Tags:  nil,
	}
	fmt.Println(product4)

	fmt.Println()

	product5 := structs.Product{
		Name: "Heladera marca sarasa",
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags: []string{"Electrodomestico", "Frio", "Cocina"},
	}

	fmt.Println(product5)

	fmt.Println()

	product6 := structs.Product{
		Name: "Heladera marca sarasa",
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Count: 5,
		Price: 20,
		Tags:  []string{"Electrodomestico", "Frio", "Cocina"},
	}

	product6JSON, err := json.Marshal(product6)

	fmt.Println(err)
	fmt.Println(product6JSON)         // esto regresa los valores en bytes
	fmt.Println(string(product6JSON)) // esto regresa los valores en bytes

	fmt.Println("Precio total: ", product6.TotalPrice())

	fmt.Println()

	fmt.Println(product6)
	product6.SetName("Segundo nombre")
	fmt.Println(product6)

	fmt.Println()
	fmt.Println("========== Car ==========")
	fmt.Println()

	p1 := structs.Product{
		Name:  "Cortina",
		Price: 2700,
		Type: structs.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"Hogar", "Cortina"},
		Count: 3,
	}

	p2 := structs.Product{
		Name:  "Mouse",
		Price: 1700,
		Type: structs.Type{
			Code:        "D",
			Description: "Electronica",
		},
		Tags:  []string{"computadoras", "gamer"},
		Count: 2,
	}

	car1 := structs.NewCar(12345)

	car1.AddProducts(p1, p2)

	fmt.Println(car1)
	fmt.Println(car1.Total())

}
