package singleton

import (
	"testing"
)

func Test_GetInstance(t *testing.T) {
	tests := []struct {
		name    string
		init    func(t *testing.T) *singleton
		inspect func(r *singleton, t *testing.T)
	}{
		{
			name: "GetInstance() not nill",
			init: func(t *testing.T) *singleton {
				return GetInstance()
			},
			inspect: func(r *singleton, t *testing.T) {
				if r == nil {
					t.Errorf("expected pointer to Singleton after calling GetInstance(), not nill")
				}
			},
		},
		{
			name: "first time count must be 1",
			init: func(t *testing.T) *singleton {
				return GetInstance()
			},
			inspect: func(r *singleton, t *testing.T) {
				currentCount := r.AddOne()
				if currentCount != 1 {
					t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
				}
			},
		},
		{
			name: "same instance in counter",
			init: func(t *testing.T) *singleton {
				return GetInstance()
			},
			inspect: func(r *singleton, t *testing.T) {
				counter := GetInstance()
				if r != counter {
					t.Errorf("Expected same instance in counter2 but it got a different instance")
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
