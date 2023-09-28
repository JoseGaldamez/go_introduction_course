package vehicleinterfaces

import "fmt"

type Vehicle interface {
	Distance() float64
}

const (
	CarVehicle       = "CAR"
	MotocicleVehicle = "MOTORCYCLE"
	TruckVehicle     = "TRUCK"
)

func New(vehicle string, time int) (Vehicle, error) {
	switch vehicle {
	case CarVehicle:
		return &Car{Time: time}, nil
	case MotocicleVehicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	}

	return nil, fmt.Errorf("Vehicle '%s' not exist.", vehicle)
}
