package collect

import (
	"context"
	"fmt"
	"math"
	"os"
	"time"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

// RunConfig holds paths and settings for a collection run.
type RunConfig struct {
	SeedPath      string
	OverridesPath string
	LogPath       string
	DomainConfigs []DomainConfig
	Collectors    []Collector
	Timeout       time.Duration
}

// Run executes the full collection pipeline for all configured domains.
func Run(cfg RunConfig) error {
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Timeout)
	defer cancel()

	seed, err := ingest.ReadSeed(cfg.SeedPath)
	if err != nil {
		return fmt.Errorf("read seed: %w", err)
	}

	overrides, err := LoadOverrides(cfg.OverridesPath)
	if err != nil {
		return fmt.Errorf("load overrides: %w", err)
	}

	// C3: Check for stale override data before proceeding.
	if warnings := CheckOverrideStaleness(overrides); len(warnings) > 0 {
		fmt.Fprintln(os.Stderr, "\n  Staleness warnings:")
		for _, w := range warnings {
			fmt.Fprintf(os.Stderr, "    ⚠ %s\n", w)
		}
		fmt.Fprintln(os.Stderr)
	}

	prevLog, err := LoadCollectLog(cfg.LogPath)
	if err != nil {
		return fmt.Errorf("load collect log: %w", err)
	}

	// Run all collectors.
	//
	// Error convention (Q6): A collector returns (results, nil) with Err set on
	// individual CollectResult entries for per-indicator failures (e.g., manual
	// stubs). A collector returns (nil, err) only for infrastructure-level
	// failures that prevent any collection attempt (e.g., context canceled).
	var allResults []CollectResult
	for _, c := range cfg.Collectors {
		fmt.Fprintf(os.Stderr, "  Collecting from %s...\n", c.Name())
		results, err := c.Collect(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "    WARN: %s failed: %v\n", c.Name(), err)
			continue
		}
		now := time.Now()
		for i := range results {
			results[i].CollectedAt = now
			if results[i].Err != nil {
				fmt.Fprintf(os.Stderr, "    MANUAL: %s → %v\n", results[i].IndicatorName, results[i].Err)
			} else {
				fmt.Fprintf(os.Stderr, "    OK: %s = %.2f %s\n", results[i].IndicatorName, results[i].RawValue, results[i].Unit)
			}
		}
		allResults = append(allResults, results...)
	}

	// Update each configured domain.
	newLog := &CollectLog{
		RunAt:   time.Now(),
		Entries: map[string][]LogEntry{},
	}

	for _, dcfg := range cfg.DomainConfigs {
		if err := updateDomain(seed, dcfg, allResults, overrides, prevLog, newLog); err != nil {
			return fmt.Errorf("update domain %s: %w", dcfg.DomainID, err)
		}
	}

	// Recompute composite score from all domain scores.
	// Always use full config set so partial runs (-domain flag) don't zero out
	// domains that weren't collected in this run.
	updateComposite(seed, AllDomainConfigs(), prevLog)

	// W5: Backup seed.json before overwriting.
	backupPath := cfg.SeedPath + ".bak"
	if data, err := os.ReadFile(cfg.SeedPath); err == nil {
		if err := os.WriteFile(backupPath, data, 0644); err != nil {
			fmt.Fprintf(os.Stderr, "  WARN: could not create backup at %s: %v\n", backupPath, err)
		}
	}

	if err := ingest.WriteSeed(cfg.SeedPath, seed); err != nil {
		return fmt.Errorf("write seed: %w", err)
	}

	if err := newLog.Save(cfg.LogPath); err != nil {
		return fmt.Errorf("write collect log: %w", err)
	}

	return nil
}

func updateDomain(
	seed *ingest.Seed,
	dcfg DomainConfig,
	results []CollectResult,
	overrides OverrideFile,
	prevLog *CollectLog,
	newLog *CollectLog,
) error {
	// Find the domain in seed.
	var domain *ingest.DomainJSON
	for i := range seed.Delegation.Domains {
		if seed.Delegation.Domains[i].ID == dcfg.DomainID {
			domain = &seed.Delegation.Domains[i]
			break
		}
	}
	if domain == nil {
		return fmt.Errorf("domain %s not found in seed", dcfg.DomainID)
	}

	// W2: Only include indicators that have a real value in the scoring map.
	// Missing indicators are excluded from the weighted average, not scored as zero.
	normalizedValues := map[string]float64{}

	for _, indCfg := range dcfg.Indicators {
		raw, unit, freshness, method, errMsg := resolveValue(
			indCfg.Name, dcfg.DomainID, results, overrides, domain.SubIndicators,
		)

		// W1: Loud warning for cached/missing fallbacks.
		switch method {
		case "cached":
			fmt.Fprintf(os.Stderr, "    ⚠ CACHED: %s — using stale value %.2f %s (no collector or override)\n",
				indCfg.Name, raw, unit)
		case "missing":
			fmt.Fprintf(os.Stderr, "    ✗ MISSING: %s — no value from any source, excluded from scoring\n",
				indCfg.Name)
		}

		// Only include in scoring if we actually have a value.
		if method != "missing" {
			normalized := Normalize(raw, indCfg.NormConfig)
			normalizedValues[indCfg.Name] = normalized
			updateSubIndicator(domain, indCfg.Name, raw, unit, indCfg.SourceName, freshness)
			fmt.Fprintf(os.Stderr, "    %s: %.2f %s [%s] → normalized %.1f\n",
				indCfg.Name, raw, unit, method, normalized)
		}

		entry := LogEntry{
			IndicatorName: indCfg.Name,
			SourceName:    indCfg.SourceName,
			Value:         raw,
			Unit:          unit,
			Freshness:     freshness,
			CollectedAt:   time.Now(),
			Method:        method,
		}
		if errMsg != "" {
			entry.Error = errMsg
		}
		newLog.Entries[dcfg.DomainID] = append(newLog.Entries[dcfg.DomainID], entry)
	}

	newScore := math.Round(ComputeDomainScore(normalizedValues, dcfg)*10) / 10

	fmt.Fprintf(os.Stderr, "  %s score: %.1f → %.1f\n", dcfg.DomainID, domain.Score, newScore)

	domain.PreviousScore = domain.Score
	domain.Score = newScore
	domain.Trend = UpdateTrend(domain.Trend, newScore, prevLog.RunAt, 84)
	domain.Status = ClassifyStatus(newScore)

	return nil
}

// resolveValue implements three-tier fallback: collected → override → cached.
// Returns method "missing" if no value is available from any source.
func resolveValue(
	name, domainID string,
	results []CollectResult,
	overrides OverrideFile,
	existingSubs []ingest.SubIndicatorJSON,
) (value float64, unit string, freshness string, method string, errMsg string) {
	// 1. Try collected results (auto).
	for _, r := range results {
		if r.IndicatorName == name && r.DomainID == domainID && r.Err == nil {
			return r.RawValue, r.Unit, r.Freshness, "auto", ""
		}
	}

	// 2. Try overrides (manual YAML).
	if entries, ok := overrides[domainID]; ok {
		for _, e := range entries {
			if e.Name == name {
				return e.Value, e.Unit, e.Freshness, "override", ""
			}
		}
	}

	// 3. Fall back to existing seed value (cached — last known good).
	for _, si := range existingSubs {
		if si.Name == name {
			return si.Value, si.Unit, si.Freshness, "cached", "no collector or override available"
		}
	}

	return 0, "", "", "missing", "no value found from any source"
}

// updateComposite recomputes the composite delegation score as the weighted
// average of all domain scores. Only domains present in both the config list
// and the seed are included.
func updateComposite(seed *ingest.Seed, configs []DomainConfig, prevLog *CollectLog) {
	// Build a weight map from configs.
	weightMap := make(map[string]float64, len(configs))
	for _, c := range configs {
		weightMap[c.DomainID] = c.Weight
	}

	var composite float64
	for _, d := range seed.Delegation.Domains {
		w, ok := weightMap[d.ID]
		if !ok {
			// Domain not in this run's config — use its existing score at its
			// known weight. This handles partial runs (-domain flag).
			continue
		}
		composite += d.Score * w
	}

	composite = math.Round(composite*10) / 10

	prev := seed.Delegation.Composite.Current
	seed.Delegation.Composite.Previous = prev
	seed.Delegation.Composite.Current = composite
	seed.Delegation.Composite.Delta = math.Round((composite-prev)*10) / 10
	seed.Delegation.Composite.Trend = UpdateTrend(
		seed.Delegation.Composite.Trend, composite, prevLog.RunAt, 84,
	)
	seed.Delegation.Composite.LastUpdated = time.Now().UTC().Format("2006-01-02")

	fmt.Fprintf(os.Stderr, "\n  Composite score: %.1f → %.1f (Δ%.1f)\n",
		prev, composite, composite-prev)
}

// updateSubIndicator updates an existing sub-indicator in the domain, or appends it.
func updateSubIndicator(domain *ingest.DomainJSON, name string, value float64, unit, source, freshness string) {
	for i := range domain.SubIndicators {
		if domain.SubIndicators[i].Name == name {
			domain.SubIndicators[i].Value = value
			domain.SubIndicators[i].Unit = unit
			domain.SubIndicators[i].Source = source
			domain.SubIndicators[i].Freshness = freshness
			return
		}
	}
	domain.SubIndicators = append(domain.SubIndicators, ingest.SubIndicatorJSON{
		Name:      name,
		Value:     value,
		Unit:      unit,
		Source:    source,
		Freshness: freshness,
	})
}
