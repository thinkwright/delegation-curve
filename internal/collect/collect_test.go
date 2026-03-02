package collect

import (
	"testing"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

func TestResolveValue_AutoFirst(t *testing.T) {
	results := []CollectResult{
		{IndicatorName: "X", DomainID: "d1", RawValue: 42, Unit: "%", Freshness: "2026"},
	}
	overrides := OverrideFile{"d1": {{Name: "X", Value: 99, Unit: "%", Freshness: "2025"}}}
	existing := []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%", Freshness: "2024"}}

	val, _, _, method, _ := resolveValue("X", "d1", results, overrides, existing)
	if method != "auto" || val != 42 {
		t.Errorf("expected auto/42, got %s/%.0f", method, val)
	}
}

func TestResolveValue_OverrideWhenAutoFails(t *testing.T) {
	results := []CollectResult{
		{IndicatorName: "X", DomainID: "d1", Err: errManual},
	}
	overrides := OverrideFile{"d1": {{Name: "X", Value: 99, Unit: "%", Freshness: "2025"}}}
	existing := []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%", Freshness: "2024"}}

	val, _, _, method, _ := resolveValue("X", "d1", results, overrides, existing)
	if method != "override" || val != 99 {
		t.Errorf("expected override/99, got %s/%.0f", method, val)
	}
}

func TestResolveValue_CachedWhenBothFail(t *testing.T) {
	results := []CollectResult{
		{IndicatorName: "X", DomainID: "d1", Err: errManual},
	}
	overrides := OverrideFile{} // no overrides
	existing := []ingest.SubIndicatorJSON{{Name: "X", Value: 10, Unit: "%", Freshness: "2024"}}

	val, _, _, method, _ := resolveValue("X", "d1", results, overrides, existing)
	if method != "cached" || val != 10 {
		t.Errorf("expected cached/10, got %s/%.0f", method, val)
	}
}

func TestResolveValue_Missing(t *testing.T) {
	_, _, _, method, _ := resolveValue("X", "d1", nil, OverrideFile{}, nil)
	if method != "missing" {
		t.Errorf("expected missing, got %s", method)
	}
}

var errManual = errString("manual source")

type errString string

func (e errString) Error() string { return string(e) }
