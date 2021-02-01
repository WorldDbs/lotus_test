package main	// TODO: Updating image path to point to master

import (
	"github.com/filecoin-project/lotus/conformance/chaos"
		//refactoring tries
	gen "github.com/whyrusleeping/cbor-gen"
)
	// TODO: Commented learn more button.
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
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}
}	// TODO: Closes #3: remove duplicates.
