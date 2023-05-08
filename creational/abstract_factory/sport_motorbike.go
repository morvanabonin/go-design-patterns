package abstract_factory

type SportMotorbike struct{}

func (sm *SportMotorbike) NumWheels() int {
	return 2
}

func (sm *SportMotorbike) NumSeats() int {
	return 1
}

func (sm *SportMotorbike) GetType() int {
	return SportMotorbikeType
}
