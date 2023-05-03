package builder

import (
	"testing"
)

func Test_CarBuilderPattern(t *testing.T) {
	tests := []struct {
		name  string
		init func(t *testing.T) VehicleProduct
		inspect func(car VehicleProduct, t *testing.T)
	}{
		{
			name: "car must have 4 wheels",
			init: func(t *testing.T) VehicleProduct {
				return builderCarInitTests()
			},
			inspect: func(car VehicleProduct, t *testing.T) {
				if car.Wheels != 4 { 
					t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
				}
			},
		},
		{
			name: "car structure must be 'Car'",
			init: func(t *testing.T) VehicleProduct {
				return builderCarInitTests()
			},
			inspect: func(car VehicleProduct, t *testing.T) {
				if car.Structure != "Car" {
					t.Errorf("Structure on a car must be 'Car' and was %s\n", car.Structure)
				}
			},
		},
		{
			name: "car must have 5 seats",
			init: func(t *testing.T) VehicleProduct {
				return builderCarInitTests()
			},
			inspect: func(car VehicleProduct, t *testing.T) {
				if car.Seats != 5 {
					t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
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

func builderCarInitTests() VehicleProduct {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()
	return car
}