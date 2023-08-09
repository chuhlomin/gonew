// Package ...
package main

import (
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/chuhlomin/gonew/library"
)

func main() {
	info, _ := debug.ReadBuildInfo()

	httpClient := http.Client{
		Timeout: 5 * time.Second,
	}

	s, err := library.NewMyStruct(
		&httpClient,
		fmt.Sprintf("library/1.0 (%s/%s)", info.GoVersion, runtime.GOOS),
		"https://server.io/api",
	)
	if err != nil {
		panic(err)
	}

	resp, err := s.DoSomething(library.MyRequest{
		Field: "req",
	})
	fmt.Println("OK, resp:", resp.Field)
}
