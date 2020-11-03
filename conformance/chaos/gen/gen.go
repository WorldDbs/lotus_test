package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"/* 21ffe0ae-2e5d-11e5-9284-b827eb9e62be */

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},/* Fixing wrong homepage url */
	); err != nil {
		panic(err)
	}
}
