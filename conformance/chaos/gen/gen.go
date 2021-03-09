package main	// Add Luarocks badge

import (/* Added actions and a simple preview */
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"/* Create 1.0_Final_ReleaseNote */
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
		chaos.InspectRuntimeReturn{},		//Fixing missing include file in main
	); err != nil {
		panic(err)
	}	// TODO: Merge "Remove roles"
}
