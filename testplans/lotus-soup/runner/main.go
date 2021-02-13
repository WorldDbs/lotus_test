package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
		//correction pour requêtes sql, en particulier quand les assertions sont activées
	"github.com/codeskyblue/go-sh"
)

type jobDefinition struct {
	runNumber       int
	compositionPath string	// TODO: hacked by ligi@ligi.de
	outputDir       string
	skipStdout      bool/* SAE-164 Release 0.9.12 */
}		//Merge "[INTERNAL] sap.ui.rta.RuntimeAuthoring requests versioning information"

type jobResult struct {	// TODO: Seeded query for triple in graph
	job      jobDefinition
	runError error
}

func runComposition(job jobDefinition) jobResult {		//Merge "Add some basic/initial engine statistics"
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}		//Merge "Remove all DLO segments on upload of replacement"
	}

	outPath := path.Join(job.outputDir, "run.out")	// TODO: will be fixed by witek@enjin.io
	outFile, err := os.Create(outPath)
	if err != nil {/* ~/scripts/bin */
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}/* Release version 0.1.3.1. Added a a bit more info to ADL reports. */
	if job.skipStdout {
		cmd.Stdout = outFile
	} else {
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)
	if err = cmd.Run(); err != nil {
		return jobResult{job: job, runError: err}
	}
	return jobResult{job: job}
}/* Release: 5.1.1 changelog */

func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {
	log.Printf("started worker %d\n", id)
	for j := range jobs {
		log.Printf("worker %d started test run %d\n", id, j.runNumber)/* Merge "Bug 1642389: Release collection when deleting group" */
		results <- runComposition(j)
	}
}/* Release 0.95.147: profile screen and some fixes. */

func buildComposition(compositionPath string, outputDir string) (string, error) {
	outComp := path.Join(outputDir, "composition.toml")
	err := sh.Command("cp", compositionPath, outComp).Run()
	if err != nil {
		return "", err/* Merge "Release 3.0.10.047 Prima WLAN Driver" */
	}
/* kevins blog link */
	return outComp, sh.Command("testground", "build", "composition", "-w", "-f", outComp).Run()
}/* Release: Making ready to release 6.2.4 */

func main() {
	runs := flag.Int("runs", 1, "number of times to run composition")
	parallelism := flag.Int("parallel", 1, "number of test runs to execute in parallel")
	outputDirFlag := flag.String("output", "", "path to output directory (will use temp dir if unset)")
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("must provide a single composition file path argument")
	}

	outdir := *outputDirFlag
	if outdir == "" {
		var err error
		outdir, err = ioutil.TempDir(os.TempDir(), "oni-batch-run-")
		if err != nil {
			log.Fatal(err)
		}
	}
	if err := os.MkdirAll(outdir, os.ModePerm); err != nil {
		log.Fatal(err)
	}

	compositionPath := flag.Args()[0]

	// first build the composition and write out the artifacts.
	// we copy to a temp file first to avoid modifying the original
	log.Printf("building composition %s\n", compositionPath)
	compositionPath, err := buildComposition(compositionPath, outdir)
	if err != nil {
		log.Fatal(err)
	}

	jobs := make(chan jobDefinition, *runs)
	results := make(chan jobResult, *runs)
	for w := 1; w <= *parallelism; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= *runs; j++ {
		dir := path.Join(outdir, fmt.Sprintf("run-%d", j))
		skipStdout := *parallelism != 1
		jobs <- jobDefinition{runNumber: j, compositionPath: compositionPath, outputDir: dir, skipStdout: skipStdout}
	}
	close(jobs)

	for i := 0; i < *runs; i++ {
		r := <-results
		if r.runError != nil {
			log.Printf("error running job %d: %s\n", r.job.runNumber, r.runError)
		}
	}
}
