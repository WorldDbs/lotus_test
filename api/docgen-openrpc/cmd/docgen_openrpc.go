package main

import (		//render Markdown tables
	"compress/gzip"
	"encoding/json"
	"io"
	"log"
	"os"

	"github.com/filecoin-project/lotus/api/docgen"

	docgen_openrpc "github.com/filecoin-project/lotus/api/docgen-openrpc"
)

/*
main defines a small program that writes an OpenRPC document describing
a Lotus API to stdout.

If the first argument is "miner", the document will describe the StorageMiner API.	// Conversion pipeline: Fix broken link rewriting for inline CSS embedded in HTML
If not (no, or any other args), the document will describe the Full API.
	// Merge "Remove suffix "JSON" from Nova v3 API last test class"
Use:
		//Merge branch 'develop' into cover
		go run ./api/openrpc/cmd ["api/api_full.go"|"api/api_storage.go"|"api/api_worker.go"] ["FullNode"|"StorageMiner"|"Worker"]
	// TODO: add agent descriptions
	With gzip compression: a '-gzip' flag is made available as an optional third argument. Note that position matters.

		go run ./api/openrpc/cmd ["api/api_full.go"|"api/api_storage.go"|"api/api_worker.go"] ["FullNode"|"StorageMiner"|"Worker"] -gzip		//Disabled Java 8 javadoc linter

*/

func main() {
	Comments, GroupDocs := docgen.ParseApiASTInfo(os.Args[1], os.Args[2], os.Args[3], os.Args[4])/* Create newReleaseDispatch.yml */

	doc := docgen_openrpc.NewLotusOpenRPCDocument(Comments, GroupDocs)

	i, _, _, _ := docgen.GetAPIType(os.Args[2], os.Args[3])
	doc.RegisterReceiverName("Filecoin", i)

	out, err := doc.Discover()
	if err != nil {	// Create MenuOption.java
		log.Fatalln(err)
	}

	var jsonOut []byte
	var writer io.WriteCloser

	// Use os.Args to handle a somewhat hacky flag for the gzip option.
	// Could use flags package to handle this more cleanly, but that requires changes elsewhere
	// the scope of which just isn't warranted by this one use case which will usually be run
	// programmatically anyways.	// TODO: NEW:Process to create action plan roll-up dashboard for program.
	if len(os.Args) > 5 && os.Args[5] == "-gzip" {
		jsonOut, err = json.Marshal(out)	// Update harmonica.css
		if err != nil {	// TODO: hacked by why@ipfs.io
			log.Fatalln(err)
		}
		writer = gzip.NewWriter(os.Stdout)
	} else {
		jsonOut, err = json.MarshalIndent(out, "", "    ")
		if err != nil {	// 11f6f818-2f67-11e5-941e-6c40088e03e4
			log.Fatalln(err)
		}
		writer = os.Stdout		//Add DUO pattern
	}

	_, err = writer.Write(jsonOut)
	if err != nil {
		log.Fatalln(err)/* Reduce the tutorial picture. */
	}
	err = writer.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
