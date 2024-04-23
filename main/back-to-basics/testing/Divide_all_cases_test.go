package testing

import "testing"

var testCases = []struct {
	name     string
	divident float64
	divisor  float64
	expected float64
	isErr    bool
}{
	{
		name: "valid-data", divident: 100.0, divisor: 10.0, expected: 10.0, isErr: false,
	},
	{
		name: "invalid-data", divident: 100.0, divisor: 0.0, expected: 0.0, isErr: true,
	},
}

func TestDivideAllCases(t *testing.T) {
	for _, tt := range testCases {
		res, err := Divide(tt.divident, tt.divisor)

		if tt.isErr {
			if err == nil {
				t.Error("Expected an error but didn't get one")
			}
		} else {
			if err != nil {
				t.Errorf("Didn't expect an error but got one: %s", err.Error())
			}
		}

		if res != tt.expected {
			t.Errorf("Expected %f but got %f", tt.expected, res)
		}
	}
}
