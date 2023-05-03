package builder

type CarBuilder struct {}

func (c *CarBuilder) SetWheels() BuildProcess {
	return nil
}

func (c *CarBuilder) SetSeats() BuildProcess {
	return nil
}

func (c *CarBuilder) SetStructure() BuildProcess {
	return nil
}

func (c *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}