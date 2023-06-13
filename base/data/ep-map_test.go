package data

import "testing"

func Example_map() {
	EP_map()
}

func TestEP_map(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example_map()
		})
	}
}
