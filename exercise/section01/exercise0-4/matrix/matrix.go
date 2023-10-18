package matrix

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Rows [][]float64
}

func (matrix Matrix) Print() {

	rows := len(matrix.Rows)
	columns := len(matrix.Rows[0])

	isCuadratica := rows == columns

	for _, row := range matrix.Rows {
		if columns != len(row) {
			isCuadratica = false
		}

		for _, columnValue := range row {
			fmt.Printf("%.2f\t", columnValue)
		}
		fmt.Println()
	}

	fmt.Println()
	fmt.Println("Es cuadratica: ", isCuadratica)
	fmt.Println()

}

func New(values ...[]float64) (Matrix, error) {

	var matrix Matrix

	columnSize := len(values[0])

	for _, row := range values {
		if columnSize != len(row) {
			return matrix, errors.New("All columns must to have the same size")
		}
		matrix.Rows = append(matrix.Rows, row)
	}

	return matrix, nil
}
