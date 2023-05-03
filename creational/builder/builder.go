package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

type ManufacturingDirector struct{}

func (f *ManufacturingDirector) Construct() {
	//Implementation hoes here
}

func (f *ManufacturingDirector) SetBuilder() {
	//Implementation goes here
}

type VehicleProduct struct {
	Wheels int
	Seats  int
	Structure string
}

type CarBuilder struct {}