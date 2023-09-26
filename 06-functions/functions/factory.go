package functions

func FactoryOperation(op Operation) func(float64, float64) float64 {
	switch op {
	case SUM:
		return sum
	case SUB:
		return sub

	case MUL:
		return mul
	case DIV:
		return div
	}

	return nil
}

func sum(x float64, y float64) float64 {
	return x + y
}
func sub(x float64, y float64) float64 {
	return x - y
}
func mul(x float64, y float64) float64 {
	return x + y
}

func div(x float64, y float64) float64 {
	if y == 0 {
		return 0
	}
	return x / y
}
