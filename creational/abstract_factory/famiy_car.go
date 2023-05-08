package abstract_factory

type FamilyCar struct{}

func (fc *FamilyCar) NumDoors() int {
	return 5
}

func (fc *FamilyCar) NumWheels() int {
	return 4
}

func (fc *FamilyCar) NumSeats() int {
	return 5
}
