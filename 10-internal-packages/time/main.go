package main

import (
	"fmt"
	"time"
)

func main() {
	showMessage := fmt.Println

	// Hora y fecha actual
	now := time.Now()
	showMessage(now)

	// la fecha que to decida
	then := time.Date(2023, 9, 15, 10, 30, 30, 0, time.UTC)
	showMessage(then)

	// sumar o restar cualquier parametro
	then = then.Add(time.Hour * 1)
	showMessage(then)

	showMessage(then.Year())
	showMessage(then.Month())
	showMessage(then.Day())
	showMessage(then.Minute())
	showMessage(then.Second())

	showMessage()
	// Comparaciones entre fechas
	showMessage("Then es antes que now:", then.Before(now))
	showMessage("Then es después que now:", then.After(now))

	// extraer la diferencia entre dos fechas
	diff := now.Sub(then)
	showMessage(diff.Hours())
	showMessage(diff.Minutes())

}
