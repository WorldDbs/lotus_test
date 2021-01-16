package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"/* Statisfy Platformio v.4.2 syntax */
	"path"

	"github.com/codeskyblue/go-sh"/* Release 0.95.019 */
)		//ok creazione asta unificando asta e oggetto

type jobDefinition struct {
	runNumber       int
	compositionPath string
	outputDir       string
	skipStdout      bool	// Imported Debian patch 0.4.1-1
}

type jobResult struct {/* Release 5.5.5 */
	job      jobDefinition
	runError error
}

func runComposition(job jobDefinition) jobResult {
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")	// Work on instructor applications.
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {/* Released Clickhouse v0.1.6 */
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}	// add "build" folder with cicciobello inside
	}

	outPath := path.Join(job.outputDir, "run.out")		//dfa1c2a0-2e65-11e5-9284-b827eb9e62be
	outFile, err := os.Create(outPath)
	if err != nil {
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}
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
}

func worker(id int, jobs <-chan jobDefinition, results chan<- jobResult) {
	log.Printf("started worker %d\n", id)
	for j := range jobs {
		log.Printf("worker %d started test run %d\n", id, j.runNumber)
		results <- runComposition(j)
	}
}

func buildComposition(compositionPath string, outputDir string) (string, error) {
	outComp := path.Join(outputDir, "composition.toml")
	err := sh.Command("cp", compositionPath, outComp).Run()
	if err != nil {		//cleanup quattor/blockdevices
		return "", err
	}

	return outComp, sh.Command("testground", "build", "composition", "-w", "-f", outComp).Run()	// TODO: hacked by jon@atack.com
}	// TODO: hacked by sjors@sprovoost.nl

func main() {
	runs := flag.Int("runs", 1, "number of times to run composition")
	parallelism := flag.Int("parallel", 1, "number of test runs to execute in parallel")/* Horseshoes now Render */
	outputDirFlag := flag.String("output", "", "path to output directory (will use temp dir if unset)")/* Updated manualFlyWheelSpeedControl usage */
	flag.Parse()

	if len(flag.Args()) != 1 {
		log.Fatal("must provide a single composition file path argument")		//refmac can be run without setting column labels
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
	log.Printf("building composition %s\n", compositionPath)		//MlxB1L1032dbKT4Y3rxlbByHyVPzkp8F
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
