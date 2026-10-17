package main

import (
	"encoding/json"
	"encoding/xml"
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
	allowUnsnapshotted := flag.Bool("allow-unsnapshotted", false, "Export seed scores that are not recorded in the current analysis run (local experiments only)")
	flag.Parse()

	// 1. Ingest
	fmt.Fprintf(os.Stderr, "Reading %s...\n", *input)
	seed, err := ingest.ReadSeed(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "Loaded %d domains\n", len(seed.Delegation.Domains))

	// Published numbers must be the current run's snapshot, and the seed's
	// history fields must match the run series they are derived from.
	if err := transform.CheckSnapshotted(seed); err != nil {
		if !*allowUnsnapshotted {
			fmt.Fprintf(os.Stderr, "Error: %v\nRecord the scores with curve-snapshot-run, or pass -allow-unsnapshotted for a local preview.\n", err)
			os.Exit(1)
		}
		fmt.Fprintf(os.Stderr, "WARNING: %v\n", err)
	}
	if err := transform.CheckDerived(seed); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\nRun curve-snapshot-run -rederive-only to rewrite them from the run series.\n", err)
		os.Exit(1)
	}
	series, hasSeries, err := transform.PublicSeries(seed)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// 2. Transform
	fmt.Fprintln(os.Stderr, "Transforming...")
	delegationRows, subRows, sourceRows := transform.Delegation(seed.Delegation.Domains)
	metaRow := transform.Meta(seed)
	runRows, domainScoreRows, indicatorRows := transform.History(seed)

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
		{"analysis_runs", func() error { return export.WriteTable[schema.AnalysisRunRow](*output, "analysis_runs", runRows) }},
		{"domain_scores", func() error {
			return export.WriteTable[schema.DomainScoreRow](*output, "domain_scores", domainScoreRows)
		}},
		{"indicator_observations", func() error {
			return export.WriteTable[schema.IndicatorObservationRow](*output, "indicator_observations", indicatorRows)
		}},
	}

	for _, t := range tables {
		if err := t.fn(); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", t.name, err)
			os.Exit(1)
		}
	}

	// Also emit meta.json for instant home page rendering (bypasses DuckDB-WASM).
	// Shape matches the frontend CompositeData interface exactly.
	delegation := map[string]any{
		"current":     metaRow.DelegationComposite,
		"previous":    metaRow.DelegationPrevious,
		"delta":       metaRow.DelegationDelta,
		"trend":       json.RawMessage(metaRow.DelegationTrend),
		"runHistory":  compositeRunHistory(series, hasSeries, runRows),
		"lastUpdated": metaRow.LastUpdated,
		"dataYear":    metaRow.DataYear,
	}
	if hasSeries {
		delegation["runId"] = series.Current.RunID
		delegation["publishedAt"] = series.Current.PublishedAt
		delegation["methodologyVersion"] = series.Current.MethodologyVersion
		if series.Previous != nil {
			delegation["previousPeriod"] = series.Previous.MeasurementPeriod
		}
	}
	metaJSON := map[string]any{
		"delegation":     delegation,
		"domains":        domainSummaries(seed),
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

	staticDir := filepath.Dir(*output)
	lastmod := metaRow.LastUpdated
	if hasSeries && series.Current.PublishedAt != "" {
		lastmod = series.Current.PublishedAt
	}
	if err := writeSitemap(filepath.Join(staticDir, "sitemap.xml"), seed, lastmod); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing sitemap.xml: %v\n", err)
		os.Exit(1)
	}

	seedBytes, err := os.ReadFile(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading seed for static copy: %v\n", err)
		os.Exit(1)
	}
	staticSeedPath := filepath.Join(staticDir, "seed.json")
	if err := os.WriteFile(staticSeedPath, seedBytes, 0o644); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing static seed.json: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintln(os.Stderr, "Done.")
}

const siteURL = "https://curve.thinkwright.ai"

// writeSitemap lists every prerendered public page, dated by the current
// run's publication so search engines see refreshes.
func writeSitemap(path string, seed *ingest.Seed, lastmod string) error {
	type url struct {
		Loc     string `xml:"loc"`
		LastMod string `xml:"lastmod,omitempty"`
	}
	type urlset struct {
		XMLName xml.Name `xml:"urlset"`
		XMLNS   string   `xml:"xmlns,attr"`
		URLs    []url    `xml:"url"`
	}
	set := urlset{XMLNS: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	add := func(p string) { set.URLs = append(set.URLs, url{Loc: siteURL + p, LastMod: lastmod}) }
	add("/")
	add("/delegation")
	for _, d := range seed.Delegation.Domains {
		add("/delegation/" + d.ID)
	}
	add("/data")
	add("/about")

	out, err := xml.MarshalIndent(set, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append([]byte(xml.Header), append(out, '\n')...), 0o644)
}

// compositeRunHistory lists the public curve points. With explicit runs it
// follows the public series up to the current run; legacy seeds fall back to
// the derived legacy run rows.
func compositeRunHistory(series transform.Series, hasSeries bool, rows []schema.AnalysisRunRow) []map[string]any {
	point := func(runID, label, publishedAt, period string, year int, method string, score float64, notes string, current bool) map[string]any {
		return map[string]any{
			"runId":              runID,
			"label":              label,
			"publishedAt":        publishedAt,
			"measurementPeriod":  period,
			"measurementYear":    year,
			"periodMidpoint":     transform.PeriodMidpoint(period, year),
			"methodologyVersion": method,
			"score":              score,
			"notes":              notes,
			"isCurrent":          current,
		}
	}

	var history []map[string]any
	if hasSeries {
		for _, r := range series.Public {
			history = append(history, point(r.RunID, r.Label, r.PublishedAt, r.MeasurementPeriod,
				r.MeasurementYear, r.MethodologyVersion, r.CompositeScore, r.Notes, r.RunID == series.Current.RunID))
		}
		return history
	}
	for _, r := range rows {
		if r.IsPublicSeries {
			history = append(history, point(r.RunID, r.Label, r.PublishedAt, r.MeasurementPeriod,
				int(r.MeasurementYear), r.MethodologyVersion, r.CompositeScore, r.Notes, r.IsCurrent))
		}
	}
	return history
}

// domainSummaries carries the headline fields of every domain so pages can
// render (and prerender) without querying Parquet.
func domainSummaries(seed *ingest.Seed) []map[string]any {
	out := make([]map[string]any, 0, len(seed.Delegation.Domains))
	for _, d := range seed.Delegation.Domains {
		out = append(out, map[string]any{
			"id":            d.ID,
			"name":          d.Name,
			"fullName":      d.FullName,
			"score":         d.Score,
			"previousScore": d.PreviousScore,
			"status":        d.Status,
			"weight":        d.Weight,
			"tier":          d.Tier,
			"description":   d.Description,
		})
	}
	return out
}
