package server

import "testing"

func TestEP_websocket(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "Start a Websocket Server Demo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			EP_websocket()
		})
	}
}
