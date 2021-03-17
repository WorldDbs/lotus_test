package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},/* Disable VS hosting process for Release builds too. */
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},	// TODO: will be fixed by jon@atack.com
		chaos.SendReturn{},
		chaos.MutateStateArgs{},	// TODO: fix endless redirect
		chaos.AbortWithArgs{},/* center site-name */
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)	// TODO: hacked by yuvalalaluf@gmail.com
	}
}
