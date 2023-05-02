package singleton

import "testing"

func Test_GetInstance(t *testing.T) {
	testes := []struct {
		name  string
		init func(t *testing.T) Singleton
		inspect func(t *testing.T)
	}{
		{
			name: "GetInstance() not nill",
			init: func(t *testing.T) Singleton {
				return GetInstance()
			},
			inspect: func(r Singleton, , t *testing.T) {
				if r == nil {
					t.Error("expected pointer to Singleton after calling GetInstance(), not nill")
				}
			}
		},
		{
			name: "First time count must be 1",
			init: func(t *testing.T) Singleton {
				return GetInstance()
			},
			inspect: func(r Singleton, , t *testing.T) {
				currentCount := r.AddOne()
				if counter != 1 {
					t.Error("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
				}
			}
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			receiver := tt.init(t)

			receiver.GetInstance()
			if tt.inspect != nil {
				tt.inspect(receiver, t)
			}

		})
	}
}