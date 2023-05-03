package builder

import (
	"testing"
)

func Test_TruckBuilderPattern(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) VehicleProduct
		inspect func(truck VehicleProduct, t *testing.T)
	}{
		{
			name: "truck must have 18 wheels",
			init: func(t *testing.T) VehicleProduct {
				return builderTruckInitTests()
			},
			inspect: func(truck VehicleProduct, t *testing.T) {
				if truck.Wheels != 18 {
					t.Errorf("Wheels on a truck must be 18 and they were %d\n", truck.Wheels)
				}
			},
		},
		{
			name: "truck structure must be 'Truck'",
			init: func(t *testing.T) VehicleProduct {
				return builderTruckInitTests()
			},
			inspect: func(truck VehicleProduct, t *testing.T) {
				if truck.Structure != "Truck" {
					t.Errorf("Structure on a truck must be 'Truck' and was %s\n", truck.Structure)
				}
			},
		},
		{
			name: "bike must have 2 seats",
			init: func(t *testing.T) VehicleProduct {
				return builderTruckInitTests()
			},
			inspect: func(truck VehicleProduct, t *testing.T) {
				if truck.Seats != 2 {
					t.Errorf("Seats on a truck must be 2 and they were %d\n", truck.Seats)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)

			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}
		})
	}
}

func builderTruckInitTests() VehicleProduct {
	manufacturingComplex := ManufacturingDirector{}
	truckBuilder := &TruckBuilder{}
	manufacturingComplex.SetBuilder(truckBuilder)
	manufacturingComplex.Construct()
	truck := truckBuilder.GetVehicle()
	return truck
}
