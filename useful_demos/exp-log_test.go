package useful_demos

import "testing"

func Example_log() {
	EP_log()
}

func TestEP_log(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Example_log()
		})
	}
}
