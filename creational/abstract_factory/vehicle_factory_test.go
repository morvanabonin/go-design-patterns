package abstract_factory

import (
	"testing"
)

func Test_CarFactory(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) Vehicle
		inspect func(c Vehicle, t *testing.T)
	}{
		{
			name: "Test Luxury Car Number of Wheels",
			init: func(t *testing.T) Vehicle {
				carF, _ := BuildFactory(CarFactoryType)
				carVehicle, _ := carF.Build(LuxuryCarType)

				return carVehicle
			},
			inspect: func(c Vehicle, t *testing.T) {
				luxuryCar, ok := c.(Car)
				if !ok {
					t.Fatal("Struct assertion has failed")
				}
				t.Logf("Luxury cars has %d doors.\n", luxuryCar.NumDoors())
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
