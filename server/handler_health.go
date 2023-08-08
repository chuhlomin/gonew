package main

import (
	"fmt"
	"net/http"
)

func (s *server) HealthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		errs := []error{}

		// check services

		if len(errs) > 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			for _, err := range errs {
				fmt.Fprintf(w, "%s\n", err.Error())
			}
			return
		}

		fmt.Fprintf(w, `OK`)
	}
}
