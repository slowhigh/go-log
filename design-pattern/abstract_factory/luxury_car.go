package abstract_factory

type LuxuryCar struct{}

func (*LuxuryCar) NumWheels() int {
	return 4
}

func (*LuxuryCar) NumSeats() int {
	return 5
}

func (*LuxuryCar) NumDoors() int {
	return 4
}