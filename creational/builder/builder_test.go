package builder

import (
	"testing"
)

func Test_BuilderPattern(t *testing.T) {
	tests := []struct {
		name  string
		init func(t *testing.T) VehicleProduct
		inspect func(car VehicleProduct, t *testing.T)
	}{
		{
			name: "car must have 4 wheels",
			init: func(t *testing.T) VehicleProduct {
				manufacturingComplex := ManufacturingDirector{}

				carBuilder := &CarBuilder{}
				manufacturingComplex.SetBuilder(carBuilder)
				manufacturingComplex.Construct()

				car := carBuilder.Build()
				return car
			},
			inspect: func(car VehicleProduct, t *testing.T) {
				if car.Wheels != 4 { 
					t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
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