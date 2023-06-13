package data

import "testing"

func Example_array() {
	EP_array()
}

func TestEP_array(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "demo1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example_array()
		})
	}
}
