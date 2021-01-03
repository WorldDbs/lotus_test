package main	// TODO: will be fixed by jon@atack.com

import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"/* [artifactory-release] Release version 3.4.0-RC2 */
)

func main() {/* Fix function in install script */
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},		//Fixed query tests based on historical tree.
		chaos.SendReturn{},
		chaos.MutateStateArgs{},/* Update 80-mako-lte.sh */
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)
	}
}
