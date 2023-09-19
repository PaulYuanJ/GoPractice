package server

import "testing"

func TestEP_http(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Start HTTP Server",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EP_http()
		})
	}
}
