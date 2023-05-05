package factory

import (
	"reflect"
	"testing"
)

func Test_CreatePaymentMethodCash(t *testing.T) {
	type args struct {
		paymentMethod int
	}
	tests := []struct {
		name    string
		args    func(t *testing.T) args
		want1   *CashPM
		wantErr bool
	}{
		{
			name: "Test method payment cash",
			args: func(t *testing.T) args {
				var a args
				a.paymentMethod = 1
				return a
			},
			want1:   new(CashPM),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := GetPaymentMethod(tArgs.paymentMethod)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPaymentMethod Cash got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetPaymentMethod Cash error = %v, wantErr: %t", err, tt.wantErr)
			}
		})
	}
}

func Test_CreatePaymentMethodDebitCard(t *testing.T) {
	type args struct {
		paymentMethod int
	}
	tests := []struct {
		name    string
		args    func(t *testing.T) args
		want1   *DebitCardPM
		wantErr bool
	}{
		{
			name: "Test method payment debit card",
			args: func(t *testing.T) args {
				var a args
				a.paymentMethod = 2
				return a
			},
			want1:   new(DebitCardPM),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := GetPaymentMethod(tArgs.paymentMethod)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPaymentMethod Debit Card got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetPaymentMethod Debit Card error = %v, wantErr: %t", err, tt.wantErr)
			}
		})
	}
}

func Test_CreatePaymentMethodCreditCard(t *testing.T) {
	type args struct {
		paymentMethod int
	}
	tests := []struct {
		name    string
		args    func(t *testing.T) args
		want1   *CreditCardPM
		wantErr bool
	}{
		{
			name: "Test method payment credit card",
			args: func(t *testing.T) args {
				var a args
				a.paymentMethod = 3
				return a
			},
			want1:   new(CreditCardPM),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := GetPaymentMethod(tArgs.paymentMethod)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetPaymentMethod Credit Card got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetPaymentMethod Credit Card error = %v, wantErr: %t", err, tt.wantErr)
			}
		})
	}
}
