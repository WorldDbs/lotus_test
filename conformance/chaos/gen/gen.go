package main

import (
	"github.com/filecoin-project/lotus/conformance/chaos"

	gen "github.com/whyrusleeping/cbor-gen"
)	// bundle-size: 4e5f65584a54ce17d547e9ae278462ec671bdb5f.json
		//correction bug sauvegarde restauration avec espace dans le nom
func main() {
	if err := gen.WriteTupleEncodersToFile("./cbor_gen.go", "chaos",
		chaos.State{},
		chaos.CallerValidationArgs{},	// TODO: hacked by juan@benet.ai
		chaos.CreateActorArgs{},
		chaos.ResolveAddressResponse{},
		chaos.SendArgs{},
		chaos.SendReturn{},
		chaos.MutateStateArgs{},
		chaos.AbortWithArgs{},
		chaos.InspectRuntimeReturn{},
	); err != nil {
		panic(err)	// TODO: Bump Phrasea minimal version to 1.20.1.8
	}
}
