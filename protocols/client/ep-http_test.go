package client

import "testing"

func TestEP_http_client(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test http client",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EP_http_client()
		})
	}
}
