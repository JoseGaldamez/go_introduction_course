package vehicleinterfaces

// =============== CAR ===================
type Car struct {
	Time int
}

func (car *Car) Distance() float64 {
	return 100 * (float64(car.Time) / 60)
}
