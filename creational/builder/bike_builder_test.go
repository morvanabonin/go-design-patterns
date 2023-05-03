package builder

import (
	"testing"
)

func Test_BikeBuilderPattern(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) VehicleProduct
		inspect func(bike VehicleProduct, t *testing.T)
	}{
		{
			name: "bike must have 2 wheels",
			init: func(t *testing.T) VehicleProduct {
				return builderBikeInitTests()
			},
			inspect: func(bike VehicleProduct, t *testing.T) {
				if bike.Wheels != 2 {
					t.Errorf("Wheels on a bike must be 2 and they were %d\n", bike.Wheels)
				}
			},
		},
		{
			name: "bike structure must be 'Motorbike'",
			init: func(t *testing.T) VehicleProduct {
				return builderBikeInitTests()
			},
			inspect: func(bike VehicleProduct, t *testing.T) {
				if bike.Structure != "Motorbike" {
					t.Errorf("Structure on a bike must be 'Motorbike' and was %s\n", bike.Structure)
				}
			},
		},
		{
			name: "bike must have 2 seats",
			init: func(t *testing.T) VehicleProduct {
				return builderBikeInitTests()
			},
			inspect: func(bike VehicleProduct, t *testing.T) {
				if bike.Seats != 2 {
					t.Errorf("Seats on a bike must be 2 and they were %d\n", bike.Seats)
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

func builderBikeInitTests() VehicleProduct {
	manufacturingComplex := ManufacturingDirector{}
	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()
	bike := bikeBuilder.GetVehicle()
	return bike
}
