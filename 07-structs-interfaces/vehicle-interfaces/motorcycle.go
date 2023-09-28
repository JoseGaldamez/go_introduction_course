package vehicleinterfaces

// =============== Motorcycle ===================
type Motorcycle struct {
	Time int
}

func (car *Motorcycle) Distance() float64 {
	return 120 * (float64(car.Time) / 60)
}
