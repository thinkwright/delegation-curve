package transform

import (
	"testing"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

func TestMetaUsesSeedDataFreshness(t *testing.T) {
	seed := &ingest.Seed{
		Delegation: ingest.DelegationSeed{
			Composite: ingest.CompositeDelegation{
				Current:       46,
				Previous:      39.7,
				Delta:         6.3,
				Trend:         []float64{39.7, 46},
				LastUpdated:   "2026-03-02",
				DataYear:      2025,
				DataFreshness: "Q4 2025",
			},
			Domains: []ingest.DomainJSON{
				{Name: "CODE-GEN", Score: 46.6},
			},
		},
	}

	got := Meta(seed)
	if got.DataFreshness != "Q4 2025" {
		t.Fatalf("DataFreshness = %q, want Q4 2025", got.DataFreshness)
	}
}

func TestMetaFallsBackToDataYear(t *testing.T) {
	seed := &ingest.Seed{
		Delegation: ingest.DelegationSeed{
			Composite: ingest.CompositeDelegation{
				Current:     46,
				Previous:    39.7,
				Delta:       6.3,
				Trend:       []float64{39.7, 46},
				LastUpdated: "2026-03-02",
				DataYear:    2025,
			},
		},
	}

	got := Meta(seed)
	if got.DataFreshness != "2025" {
		t.Fatalf("DataFreshness = %q, want 2025", got.DataFreshness)
	}
}
