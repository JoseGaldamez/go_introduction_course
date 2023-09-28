package vehicleinterfaces

// =============== Truck ===================
type Truck struct {
	Time int
}

func (car *Truck) Distance() float64 {
	return 70 * (float64(car.Time) / 60)
}
