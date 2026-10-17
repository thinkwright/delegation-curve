package collect

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

func TestResolveValue_AutoFirst(t *testing.T) {
	results := []CollectResult{
		{IndicatorName: "X", DomainID: "d1", RawValue: 42, Unit: "%", SourceName: "Auto Source", Freshness: "2026"},
	}
	overrides := OverrideFile{"d1": {{Name: "X", Value: 99, Unit: "%", Source: "Override Source", Freshness: "2025"}}}
	existing := []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%", Source: "Cached Source", Freshness: "2024"}}

	r := resolveValue("X", "d1", results, overrides, existing)
	if r.Method != "auto" || r.Value != 42 || r.Source != "Auto Source" {
		t.Errorf("expected auto/42/Auto Source, got %s/%.0f/%s", r.Method, r.Value, r.Source)
	}
}

func TestResolveValue_OverrideWhenAutoFails(t *testing.T) {
	results := []CollectResult{
		{IndicatorName: "X", DomainID: "d1", Err: errManual},
	}
	overrides := OverrideFile{"d1": {{Name: "X", Value: 99, Unit: "%", Source: "Override Source", Freshness: "2025"}}}
	existing := []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%", Source: "Cached Source", Freshness: "2024"}}

	r := resolveValue("X", "d1", results, overrides, existing)
	if r.Method != "override" || r.Value != 99 || r.Source != "Override Source" {
		t.Errorf("expected override/99/Override Source, got %s/%.0f/%s", r.Method, r.Value, r.Source)
	}
}

func TestResolveValue_CachedWhenBothFail(t *testing.T) {
	results := []CollectResult{
		{IndicatorName: "X", DomainID: "d1", Err: errManual},
	}
	overrides := OverrideFile{} // no overrides
	existing := []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%", Source: "Cached Source", Freshness: "2024"}}

	r := resolveValue("X", "d1", results, overrides, existing)
	if r.Method != "cached" || r.Value != 10 || r.Source != "Cached Source" {
		t.Errorf("expected cached/10/Cached Source, got %s/%.0f/%s", r.Method, r.Value, r.Source)
	}
}

func TestResolveValue_Missing(t *testing.T) {
	if r := resolveValue("X", "d1", nil, OverrideFile{}, nil); r.Method != "missing" {
		t.Errorf("expected missing, got %s", r.Method)
	}
}

func TestUpdateComposite_DoesNotTouchHistory(t *testing.T) {
	seed := &ingest.Seed{
		Delegation: ingest.DelegationSeed{
			Composite: ingest.CompositeDelegation{
				Current:  45,
				Previous: 40,
				Delta:    5,
				Trend:    []float64{40, 45},
			},
			Domains: []ingest.DomainJSON{
				{ID: "d1", Score: 50},
				{ID: "d2", Score: 30},
			},
		},
	}
	configs := []DomainConfig{{DomainID: "d1", Weight: 0.75}, {DomainID: "d2", Weight: 0.25}}

	updateComposite(seed, configs)

	c := seed.Delegation.Composite
	if c.Current != 45 {
		t.Fatalf("Current = %v, want 45", c.Current)
	}
	if c.Previous != 40 || c.Delta != 5 || len(c.Trend) != 2 {
		t.Fatalf("history fields changed: %+v", c)
	}
}

func TestUpdateComposite_NormalizesByPresentWeight(t *testing.T) {
	seed := &ingest.Seed{Delegation: ingest.DelegationSeed{Domains: []ingest.DomainJSON{{ID: "d1", Score: 50}}}}
	updateComposite(seed, []DomainConfig{{DomainID: "d1", Weight: 0.5}, {DomainID: "absent", Weight: 0.5}})
	if seed.Delegation.Composite.Current != 50 {
		t.Fatalf("Current = %v, want 50", seed.Delegation.Composite.Current)
	}
}

func TestRunDryRunWritesNothing(t *testing.T) {
	dir := t.TempDir()
	seedPath := filepath.Join(dir, "seed.json")
	logPath := filepath.Join(dir, "collect.log.json")
	seed := &ingest.Seed{Delegation: ingest.DelegationSeed{
		Composite: ingest.CompositeDelegation{Current: 10, Previous: 8, Delta: 2, Trend: []float64{8, 10}},
		Domains: []ingest.DomainJSON{{
			ID: "d1", Score: 10, PreviousScore: 8, Trend: []float64{8, 10},
			SubIndicators: []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%"}},
		}},
	}}
	if err := ingest.WriteSeed(seedPath, seed); err != nil {
		t.Fatal(err)
	}
	before, _ := os.ReadFile(seedPath)

	cfg := RunConfig{
		SeedPath:         seedPath,
		OverridesPath:    filepath.Join(dir, "missing-overrides.yaml"),
		LogPath:          logPath,
		DomainConfigs:    []DomainConfig{{DomainID: "d1", Weight: 1, Indicators: []IndicatorConfig{{Name: "X", Weight: 1}}}},
		CompositeConfigs: []DomainConfig{{DomainID: "d1", Weight: 1}},
		Timeout:          time.Second,
		DryRun:           true,
		Check:            true,
	}
	if err := Run(cfg); err != nil {
		t.Fatalf("unchanged checked dry run returned %v", err)
	}
	after, _ := os.ReadFile(seedPath)
	if string(before) != string(after) {
		t.Fatal("dry run modified the seed")
	}
	if _, err := os.Stat(logPath); !os.IsNotExist(err) {
		t.Fatal("dry run wrote the collect log")
	}

	// A changed override must fail a checked dry run.
	overrides := filepath.Join(dir, "overrides.yaml")
	os.WriteFile(overrides, []byte("d1:\n  - name: X\n    value: 20\n    unit: \"%\"\n"), 0o644)
	cfg.OverridesPath = overrides
	if err := Run(cfg); !errors.Is(err, ErrScoresChanged) {
		t.Fatalf("changed checked dry run returned %v, want ErrScoresChanged", err)
	}
}

func TestRunKeepsHistoryFields(t *testing.T) {
	dir := t.TempDir()
	seedPath := filepath.Join(dir, "seed.json")
	seed := &ingest.Seed{Delegation: ingest.DelegationSeed{
		Composite: ingest.CompositeDelegation{Current: 10, Previous: 8, Delta: 2, Trend: []float64{8, 10}},
		Domains: []ingest.DomainJSON{{
			ID: "d1", Score: 10, PreviousScore: 8, Trend: []float64{8, 10},
			SubIndicators: []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%"}},
		}},
	}}
	if err := ingest.WriteSeed(seedPath, seed); err != nil {
		t.Fatal(err)
	}
	overrides := filepath.Join(dir, "overrides.yaml")
	os.WriteFile(overrides, []byte("d1:\n  - name: X\n    value: 20\n    unit: \"%\"\n"), 0o644)

	err := Run(RunConfig{
		SeedPath: seedPath, OverridesPath: overrides, LogPath: filepath.Join(dir, "log.json"),
		DomainConfigs:    []DomainConfig{{DomainID: "d1", Weight: 1, Indicators: []IndicatorConfig{{Name: "X", Weight: 1}}}},
		CompositeConfigs: []DomainConfig{{DomainID: "d1", Weight: 1}},
		Timeout:          time.Second,
	})
	if err != nil {
		t.Fatal(err)
	}
	got, err := ingest.ReadSeed(seedPath)
	if err != nil {
		t.Fatal(err)
	}
	d := got.Delegation.Domains[0]
	if d.Score != 20 || got.Delegation.Composite.Current != 20 {
		t.Fatalf("scores not updated: domain %v composite %v", d.Score, got.Delegation.Composite.Current)
	}
	if d.PreviousScore != 8 || len(d.Trend) != 2 || got.Delegation.Composite.Previous != 8 || len(got.Delegation.Composite.Trend) != 2 {
		t.Fatalf("collect rewrote history: domain %+v composite %+v", d, got.Delegation.Composite)
	}
}

var errManual = errString("manual source")

type errString string

func (e errString) Error() string { return string(e) }
