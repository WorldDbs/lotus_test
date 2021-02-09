package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)
/* Create Elli.json */
func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},/* add ADC port defines in NanoRelease1.h, this pin is used to pull the Key pin */
		chaos.SendArgs{},
		chaos.SendReturn{},/* Merge "Migrate synchronizer to DSE2" */
		chaos.MutateStateArgs{},		//Add handler variable
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}
}
