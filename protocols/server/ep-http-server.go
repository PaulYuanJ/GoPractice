package server

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	fmt.Fprintf(w, "Hello World %s with headers: %v", "Paul", w.Header())
}

func EP_http() {
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: "0.0.0.0:8099",
	}
	_ = server.ListenAndServe()
}
