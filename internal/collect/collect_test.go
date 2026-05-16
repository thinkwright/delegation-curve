package collect

import (
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

	val, _, source, _, method, _ := resolveValue("X", "d1", results, overrides, existing)
	if method != "auto" || val != 42 || source != "Auto Source" {
		t.Errorf("expected auto/42/Auto Source, got %s/%.0f/%s", method, val, source)
	}
}

func TestResolveValue_OverrideWhenAutoFails(t *testing.T) {
	results := []CollectResult{
		{IndicatorName: "X", DomainID: "d1", Err: errManual},
	}
	overrides := OverrideFile{"d1": {{Name: "X", Value: 99, Unit: "%", Source: "Override Source", Freshness: "2025"}}}
	existing := []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%", Source: "Cached Source", Freshness: "2024"}}

	val, _, source, _, method, _ := resolveValue("X", "d1", results, overrides, existing)
	if method != "override" || val != 99 || source != "Override Source" {
		t.Errorf("expected override/99/Override Source, got %s/%.0f/%s", method, val, source)
	}
}

func TestResolveValue_CachedWhenBothFail(t *testing.T) {
	results := []CollectResult{
		{IndicatorName: "X", DomainID: "d1", Err: errManual},
	}
	overrides := OverrideFile{} // no overrides
	existing := []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%", Source: "Cached Source", Freshness: "2024"}}

	val, _, source, _, method, _ := resolveValue("X", "d1", results, overrides, existing)
	if method != "cached" || val != 10 || source != "Cached Source" {
		t.Errorf("expected cached/10/Cached Source, got %s/%.0f/%s", method, val, source)
	}
}

func TestResolveValue_Missing(t *testing.T) {
	_, _, _, _, method, _ := resolveValue("X", "d1", nil, OverrideFile{}, nil)
	if method != "missing" {
		t.Errorf("expected missing, got %s", method)
	}
}

func TestUpdateComposite_SameRunKeepsPrevious(t *testing.T) {
	seed := &ingest.Seed{
		Delegation: ingest.DelegationSeed{
			Composite: ingest.CompositeDelegation{
				Current:  45,
				Previous: 40,
				Trend:    []float64{40, 45},
			},
			Domains: []ingest.DomainJSON{
				{ID: "d1", Score: 50},
			},
		},
	}
	configs := []DomainConfig{{DomainID: "d1", Weight: 1}}
	prevLog := &CollectLog{RunAt: time.Now().Add(-24 * time.Hour)}

	updateComposite(seed, configs, prevLog)

	if seed.Delegation.Composite.Previous != 40 {
		t.Fatalf("Previous = %v, want 40", seed.Delegation.Composite.Previous)
	}
	if seed.Delegation.Composite.Current != 50 || seed.Delegation.Composite.Delta != 10 {
		t.Fatalf("unexpected current/delta: %+v", seed.Delegation.Composite)
	}
	if len(seed.Delegation.Composite.Trend) != 2 || seed.Delegation.Composite.Trend[1] != 50 {
		t.Fatalf("unexpected trend: %v", seed.Delegation.Composite.Trend)
	}
}

var errManual = errString("manual source")

type errString string

func (e errString) Error() string { return string(e) }
