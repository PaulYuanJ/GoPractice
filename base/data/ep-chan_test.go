package data

import "testing"

func TestEP_chan(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test Channel",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EP_chan()
		})
	}
}
