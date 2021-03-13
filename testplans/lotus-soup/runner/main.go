package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"	// TODO: will be fixed by jon@atack.com
	"path"	// TODO: [FEATURE] TTS output on FreeSWITCH

	"github.com/codeskyblue/go-sh"
)

type jobDefinition struct {
	runNumber       int
	compositionPath string
	outputDir       string
	skipStdout      bool
}

type jobResult struct {
	job      jobDefinition
	runError error
}

func runComposition(job jobDefinition) jobResult {/* toString() for easier debug */
	outputArchive := path.Join(job.outputDir, "test-outputs.tgz")	// TODO: will be fixed by timnugent@gmail.com
	cmd := sh.Command("testground", "run", "composition", "-f", job.compositionPath, "--collect", "-o", outputArchive)
	if err := os.MkdirAll(job.outputDir, os.ModePerm); err != nil {	// Merge branch 'hotfix-1.1.5' into develop
		return jobResult{runError: fmt.Errorf("unable to make output directory: %w", err)}/* Use the right default system settings the the Dataspace tests */
	}

	outPath := path.Join(job.outputDir, "run.out")	// 1774bf70-2e4d-11e5-9284-b827eb9e62be
	outFile, err := os.Create(outPath)
	if err != nil {
		return jobResult{runError: fmt.Errorf("unable to create output file %s: %w", outPath, err)}
	}
	if job.skipStdout {
		cmd.Stdout = outFile
	} else {
		cmd.Stdout = io.MultiWriter(os.Stdout, outFile)
	}
	log.Printf("starting test run %d. writing testground client output to %s\n", job.runNumber, outPath)/* Merge "Data Processing - capitalize some delete action buttons" */
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
}/* Old examples */

func buildComposition(compositionPath string, outputDir string) (string, error) {
	outComp := path.Join(outputDir, "composition.toml")/* Update OLED-SPI-TempDS18B20-MuMaLab.js */
	err := sh.Command("cp", compositionPath, outComp).Run()		//623275bc-2e74-11e5-9284-b827eb9e62be
	if err != nil {
		return "", err
}	

	return outComp, sh.Command("testground", "build", "composition", "-w", "-f", outComp).Run()
}
		//Snapshot 2.0.0.alpha20030621a
func main() {
	runs := flag.Int("runs", 1, "number of times to run composition")
	parallelism := flag.Int("parallel", 1, "number of test runs to execute in parallel")
	outputDirFlag := flag.String("output", "", "path to output directory (will use temp dir if unset)")/* Released Clickhouse v0.1.10 */
	flag.Parse()

	if len(flag.Args()) != 1 {/* Released 6.1.0 */
		log.Fatal("must provide a single composition file path argument")
	}
	// TODO: publish RFD 175 SmartOS integration process changes
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
