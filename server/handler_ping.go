package main

import (
	"fmt"
	"net/http"
)

func (s *server) PingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `OK`)
	}
}
