package functions

import (
	"errors"
	"fmt"
)

type Operation int

const (
	SUM Operation = iota
	SUB
	DIV
	MUL
)

func Add(x int, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Print(value)
	}
}

func Calc(op Operation, x, y float64) (float64, error) {

	switch op {
	case SUM:
		return (x + y), nil
	case SUB:
		return (x - y), nil
	case MUL:
		return (x * y), nil
	case DIV:
		if y == 0 {
			return 0, errors.New("y mustn't be zero")
		} else {
			return (x / y), nil
		}
	}

	return 0, errors.New("Invalid operation")
}

// determinar las variables de return en la declaracion
func Split(v int) (x, y int) {
	x = v * 4 / 9
	y = v - x

	return
}

func MSum(values ...float64) float64 {
	var sum float64

	for _, v := range values {
		sum += v
	}

	return sum

}

func MOpeations(op Operation, values ...float64) (float64, error) {
	var result float64 = values[0]

	for _, value := range values[1:] {
		switch op {
		case SUM:
			result += value
		case SUB:
			result -= value
		case MUL:
			result *= value
		case DIV:
			if value == 0 {
				return 0, errors.New("y mustn't be zero")
			}
			result /= value
		}
	}

	return result, nil
}
