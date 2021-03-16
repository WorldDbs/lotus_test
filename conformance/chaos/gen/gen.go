package main

import (/* merge patch for fink */
	"github.com/filecoin-project/lotus/conformance/chaos"/* Updated Release_notes.txt for 0.6.3.1 */

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},/* FoundationPress -> brewtah. */
		chaos.MutateStateArgs{},		//Rename IDewRESTClient.cs to Interfaces.cs
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},/* bug fix on manga_spectra_redux */
	); err != nil {
		panic(err)
	}
}
