package builder

type TruckBuilder struct {
	v VehicleProduct
}

func (t *TruckBuilder) SetWheels() BuildProcess {
	t.v.Wheels = 6*3
	return t
}

func (t *TruckBuilder) SetSeats() BuildProcess {
	t.v.Seats = 2
	return t
}

func (t *TruckBuilder) SetStructure() BuildProcess {
	t.v.Structure = "Truck"
	return t
}

func (t *TruckBuilder) GetVehicle() VehicleProduct {
	return t.v
}