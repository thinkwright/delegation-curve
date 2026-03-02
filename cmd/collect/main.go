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
	flag.Parse()

	fmt.Fprintln(os.Stderr, "AHI Data Collection")
	fmt.Fprintln(os.Stderr, "===================")

	domainConfigs := collect.AllDomainConfigs()

	allCollectors := []collect.Collector{
		// content-mod
		collectors.NewContentModCollector(),
		// algo-trade
		collectors.NewAlgoTradeCollector(),
		// code-gen
		collectors.NewVSCodeCollector(),
		collectors.NewStackOverflowCollector(),
		collectors.NewOctoverseCollector(),
		collectors.NewGitClearCollector(),
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
	}

	if err := collect.Run(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "\nDone. Run 'make generate' to rebuild Parquet files.")
}
