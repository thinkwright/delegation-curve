// Command curve-diagnostics reports what a run's composite is made of:
// composition by construct and evidence grade, data age, sensitivity to
// alternative aggregations, an uncertainty interval, and the methodology v3
// depth/reach shadow index. It never modifies the seed.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/thinkwright/delegation-curve/internal/diagnostics"
	"github.com/thinkwright/delegation-curve/internal/ingest"
	"github.com/thinkwright/delegation-curve/internal/transform"
)

func main() {
	seedPath := flag.String("seed", "seed/seed.json", "Path to seed.json")
	runID := flag.String("run-id", "", "Run to analyze (default: the current run)")
	flag.Parse()

	seed, err := ingest.ReadSeed(*seedPath)
	if err != nil {
		fail(err)
	}
	series, ok, err := transform.PublicSeries(seed)
	if err != nil {
		fail(err)
	}
	if !ok {
		fail(fmt.Errorf("seed has no explicit analysis runs"))
	}

	run := series.Current
	if *runID != "" {
		run = nil
		for i := range seed.AnalysisRuns {
			if seed.AnalysisRuns[i].RunID == *runID {
				run = &seed.AnalysisRuns[i]
			}
		}
		if run == nil {
			fail(fmt.Errorf("run %q not found", *runID))
		}
	}

	asOf, err := time.Parse("2006-01-02", run.PublishedAt)
	if err != nil {
		asOf = time.Now().UTC()
	}
	inputs, err := diagnostics.Inputs(seed, run.RunID)
	if err != nil {
		fail(err)
	}
	fmt.Print(diagnostics.Analyze(run.RunID, inputs, run.CompositeScore, asOf).Markdown())
}

func fail(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
