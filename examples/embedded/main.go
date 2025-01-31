package main

import (
	"fmt"
	"log"

	"github.com/suborbital/sat/sat"
	"github.com/suborbital/sat/sat/metrics"
)

// If you are using a control plane server, set the following environment variables before running
// 	SAT_CONTROL_PLANE={Control Plane URL}
// 	SAT_ENV_TOKEN={Environment token}
//
// If you are using a control plane server, pass the FQFN as the Runnable Arg (like below)
// See also https://docs.suborbital.dev/compute/concepts/fully-qualified-function-names
//
// If you are NOT using a control plane server, pass the path to the .wasm file on disk you'd like to load
//
func main() {
	config, _ := sat.ConfigFromRunnableArg("com.suborbital.acmeco#default::embed@v1.0.0")

	s, _ := sat.New(config, nil, metrics.SetupNoopMetrics())

	for i := 1; i < 100; i++ {
		resp, err := s.Exec([]byte("world!"))
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s\n", resp.Output)
	}
}
