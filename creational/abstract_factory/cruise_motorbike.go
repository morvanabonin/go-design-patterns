package abstract_factory

type CruiseMotorbike struct{}

func (cm *CruiseMotorbike) NumWheels() int {
	return 2
}

func (cm *CruiseMotorbike) NumSeats() int {
	return 2
}

func (cm *CruiseMotorbike) GetType() int {
	return CruiseMotorbikeType
}
