package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/thinkwright/delegation-curve/internal/export"
	"github.com/thinkwright/delegation-curve/internal/ingest"
	"github.com/thinkwright/delegation-curve/internal/schema"
	"github.com/thinkwright/delegation-curve/internal/transform"
)

func main() {
	input := flag.String("input", "seed/seed.json", "Path to seed.json")
	output := flag.String("output", "frontend/static/data", "Output directory for Parquet files")
	flag.Parse()

	// 1. Ingest
	fmt.Fprintf(os.Stderr, "Reading %s...\n", *input)
	seed, err := ingest.ReadSeed(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Loaded %d domains\n", len(seed.Delegation.Domains))

	// 2. Transform
	fmt.Fprintln(os.Stderr, "Transforming...")
	delegationRows, subRows, sourceRows := transform.Delegation(seed.Delegation.Domains)
	metaRow := transform.Meta(seed)

	// 3. Export
	fmt.Fprintf(os.Stderr, "Writing Parquet to %s/\n", *output)
	tables := []struct {
		name string
		fn   func() error
	}{
		{"delegation", func() error { return export.WriteTable[schema.DelegationRow](*output, "delegation", delegationRows) }},
		{"sub_indicators", func() error { return export.WriteTable[schema.SubIndicatorRow](*output, "sub_indicators", subRows) }},
		{"data_sources", func() error { return export.WriteTable[schema.DataSourceRow](*output, "data_sources", sourceRows) }},
		{"meta", func() error { return export.WriteTable[schema.MetaRow](*output, "meta", []schema.MetaRow{metaRow}) }},
	}

	for _, t := range tables {
		if err := t.fn(); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", t.name, err)
			os.Exit(1)
		}
	}

	// Also emit meta.json for instant home page rendering (bypasses DuckDB-WASM).
	// Shape matches the frontend CompositeData interface exactly.
	metaJSON := map[string]any{
		"delegation": map[string]any{
			"current":     metaRow.DelegationComposite,
			"previous":    metaRow.DelegationPrevious,
			"delta":       metaRow.DelegationDelta,
			"trend":       json.RawMessage(metaRow.DelegationTrend),
			"lastUpdated": metaRow.LastUpdated,
			"dataYear":    metaRow.DataYear,
		},
		"domainsTracked": metaRow.DomainsTracked,
		"highestDomain": map[string]any{
			"name":  metaRow.HighestDomainName,
			"score": metaRow.HighestDomainScore,
		},
		"dataFreshness": metaRow.DataFreshness,
	}
	jsonBytes, err := json.Marshal(metaJSON)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling meta.json: %v\n", err)
		os.Exit(1)
	}
	if err := os.WriteFile(filepath.Join(*output, "meta.json"), jsonBytes, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing meta.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Done.")
}
