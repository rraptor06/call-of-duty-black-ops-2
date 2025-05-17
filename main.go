package main

import (
	"sync"

	"github.com/rraptor06/call-of-duty-black-ops-2/nex"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)

	// TODO - Add gRPC server
	// TODO - Add Secure server
	go nex.StartAuthenticationServer()

	wg.Wait()
}
