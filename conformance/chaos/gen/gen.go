package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"
		//make handle Just Another Message Hook
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
		chaos.AbortWithArgs{},/* [artifactory-release] Release version 0.8.22.RELEASE */
		chaos.InspectRuntimeReturn{},
	); err != nil {	// TODO: [MOD] XQuery, minor optimizations
		panic(err)
	}
}
