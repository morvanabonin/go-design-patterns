package abstract_factory

type LuxuryCar struct{}

func (lc *LuxuryCar) NumDoors() int {
	return 4
}

func (lc *LuxuryCar) NumWheels() int {
	return 4
}

func (lc *LuxuryCar) NumSeats() int {
	return 5
}