package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/thinkwright/delegation-curve/internal/collect"
	"github.com/thinkwright/delegation-curve/internal/collect/collectors"
)

func main() {
	seedPath := flag.String("seed", "seed/seed.json", "Path to seed.json")
	overridesPath := flag.String("overrides", "seed/overrides.yaml", "Path to overrides.yaml")
	logPath := flag.String("log", "seed/collect.log.json", "Path to collect log")
	timeout := flag.Duration("timeout", 30*time.Second, "Collection timeout")
	domain := flag.String("domain", "", "Collect only this domain (empty = all)")
	dryRun := flag.Bool("dry-run", false, "Report value and score changes without writing the seed or log")
	offline := flag.Bool("offline", false, "Skip network collectors; resolve values from overrides and cached seed values only")
	check := flag.Bool("check", false, "With -dry-run: exit 1 if any value or score would change (reproducibility check)")
	flag.Parse()

	if *check && !*dryRun {
		fmt.Fprintln(os.Stderr, "Error: -check requires -dry-run")
		os.Exit(2)
	}

	fmt.Fprintln(os.Stderr, "Delegation Curve data collection")
	fmt.Fprintln(os.Stderr, "================================")

	domainConfigs := collect.AllDomainConfigs()

	allCollectors := []collect.Collector{
		// content-mod
		collectors.NewContentModCollector(),
		// algo-trade
		collectors.NewAlgoTradeCollector(),
		// code-gen
		collectors.NewVSCodeCollector(),
		// support
		collectors.NewSupportCollector(),
		// credit
		collectors.NewCreditCollector(),
		// medical-dx
		collectors.NewMedicalDxCollector(),
		// legal-ai
		collectors.NewLegalAICollector(),
		// hire
		collectors.NewHireCollector(),
		// education
		collectors.NewEducationCollector(),
	}

	if *domain != "" {
		var filteredD []collect.DomainConfig
		for _, d := range domainConfigs {
			if d.DomainID == *domain {
				filteredD = append(filteredD, d)
			}
		}
		domainConfigs = filteredD

		var filteredC []collect.Collector
		for _, c := range allCollectors {
			if c.DomainID() == *domain {
				filteredC = append(filteredC, c)
			}
		}
		allCollectors = filteredC
	}

	if *offline {
		allCollectors = nil
	}

	if len(domainConfigs) == 0 {
		fmt.Fprintf(os.Stderr, "Error: no domain config found for %q\n", *domain)
		os.Exit(1)
	}

	cfg := collect.RunConfig{
		SeedPath:      *seedPath,
		OverridesPath: *overridesPath,
		LogPath:       *logPath,
		DomainConfigs: domainConfigs,
		Collectors:    allCollectors,
		Timeout:       *timeout,
		DryRun:        *dryRun,
		Check:         *check,
	}

	if err := collect.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if !*dryRun {
		fmt.Fprintln(os.Stderr, "\nDone. Record the scores with curve-snapshot-run, then run 'make generate'.")
	}
}
