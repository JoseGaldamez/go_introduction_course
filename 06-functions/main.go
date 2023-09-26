package main

import (
	"fmt"

	"github.com/JoseGaldamez/go_introduction_course/06-functions/functions"
)

func main() {
	fmt.Println("Result:", functions.Add(3, 4))
	functions.RepeatString(10, "As")

	fmt.Println()

	result, err := functions.Calc(functions.SUM, 5, 5)
	fmt.Println("\nSUM value: ", result, " - error: ", err)

	result, err = functions.Calc(functions.SUB, 20, 5)
	fmt.Println("SUB value: ", result, " - error: ", err)

	result, err = functions.Calc(functions.MUL, 20, 5)
	fmt.Println("MUL value: ", result, " - error: ", err)

	result, err = functions.Calc(functions.DIV, 20, 5)
	fmt.Println("DIV value: ", result, " - error: ", err)

	result, err = functions.Calc(functions.DIV, 20, 0)
	fmt.Println("DIV value: ", result, " - error: ", err)

	fmt.Println()

	x, y := functions.Split(20)
	fmt.Println("x: ", x, " y:", y)

	fmt.Println()

	resultMultiSum := functions.MSum(3, 4, 5)
	fmt.Println("total: ", resultMultiSum)

	fmt.Println()

	resultMultiOperations, e := functions.MOpeations(functions.SUM, 10, 2, 2, 2)
	fmt.Println(resultMultiOperations, e)
	resultMultiOperations, e = functions.MOpeations(functions.SUB, 10, 2, 2, 2)
	fmt.Println(resultMultiOperations, e)
	resultMultiOperations, e = functions.MOpeations(functions.MUL, 10, 2, 2, 2)
	fmt.Println(resultMultiOperations, e)
	resultMultiOperations, e = functions.MOpeations(functions.DIV, 10, 2, 2, 2)
	fmt.Println(resultMultiOperations, e)
	resultMultiOperations, e = functions.MOpeations(functions.DIV, 10, 2, 0, 2)
	fmt.Println(resultMultiOperations, e)

	fmt.Println()

	fnSum := functions.FactoryOperation(functions.SUM)
	fmt.Println(fnSum(2, 3))
	fnSum = functions.FactoryOperation(functions.SUB)
	fmt.Println(fnSum(2, 3))
	fnSum = functions.FactoryOperation(functions.MUL)
	fmt.Println(fnSum(2, 3))
	fnSum = functions.FactoryOperation(functions.DIV)
	fmt.Println(fnSum(2, 3))

}
