package transform

import (
	"testing"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

func TestHistoryDerivesLegacyRuns(t *testing.T) {
	seed := &ingest.Seed{
		Delegation: ingest.DelegationSeed{
			Composite: ingest.CompositeDelegation{
				Trend:       []float64{39.7, 46},
				LastUpdated: "2026-03-02",
				DataYear:    2025,
			},
			Domains: []ingest.DomainJSON{
				{
					ID:     "code-gen",
					Weight: 0.15,
					Trend:  []float64{44, 52.3},
					SubIndicators: []ingest.SubIndicatorJSON{
						{
							Name:      "AI-Generated or Assisted Committed Code",
							Value:     42,
							Unit:      "%",
							Source:    "Sonar State of Code Developer Survey",
							Freshness: "2026",
						},
					},
				},
			},
		},
	}

	runs, domainScores, observations := History(seed)

	if len(runs) != 2 {
		t.Fatalf("expected 2 runs, got %d", len(runs))
	}
	if runs[0].RunID != "legacy-2024" || runs[0].MeasurementYear != 2024 || runs[0].IsCurrent {
		t.Fatalf("unexpected first run: %+v", runs[0])
	}
	if runs[1].RunID != "legacy-2025" || runs[1].PublishedAt != "2026-03-02" || !runs[1].IsCurrent {
		t.Fatalf("unexpected current run: %+v", runs[1])
	}

	if len(domainScores) != 2 {
		t.Fatalf("expected 2 domain scores, got %d", len(domainScores))
	}
	if domainScores[1].RunID != "legacy-2025" || domainScores[1].DomainID != "code-gen" || domainScores[1].Score != 52.3 {
		t.Fatalf("unexpected current domain score: %+v", domainScores[1])
	}

	if len(observations) != 1 {
		t.Fatalf("expected 1 observation, got %d", len(observations))
	}
	got := observations[0]
	if got.RunID != "legacy-2025" || !got.IncludedInScore || got.Weight != 0.45 || got.NormalizedValue != 42 {
		t.Fatalf("unexpected current observation: %+v", got)
	}
}

func TestHistoryUsesExplicitRuns(t *testing.T) {
	seed := &ingest.Seed{
		AnalysisRuns: []ingest.AnalysisRunJSON{
			{
				RunID:              "legacy-2025",
				Label:              "2025 published baseline",
				PublishedAt:        "2026-03-02",
				MeasurementPeriod:  "2025",
				MeasurementYear:    2025,
				MethodologyVersion: "legacy-trend-v1",
				CompositeScore:     46,
				Notes:              "Archived run.",
				IsPublicSeries:     boolPtr(false),
			},
			{
				RunID:              "2026-q2",
				Label:              "2026 Q2 refresh",
				PublishedAt:        "2026-06-30",
				MeasurementPeriod:  "2026 Q2",
				MeasurementYear:    2026,
				MethodologyVersion: "delegation-curve-v2",
				CompositeScore:     51.2,
				Notes:              "Explicit run.",
				IsCurrent:          true,
			},
		},
		DomainScores: []ingest.DomainScoreJSON{
			{
				RunID:    "2026-q2",
				DomainID: "code-gen",
				Score:    58.4,
				Weight:   0.15,
				Status:   "elevated",
			},
		},
		IndicatorObservations: []ingest.IndicatorObservationJSON{
			{
				RunID:           "2026-q2",
				DomainID:        "code-gen",
				IndicatorName:   "AI usage survey",
				RawValue:        62,
				NormalizedValue: 62,
				Unit:            "%",
				Weight:          0.2,
				IncludedInScore: true,
				Source:          "METR",
				SourceURL:       "https://metr.org/blog/2026-05-11-ai-usage-survey/",
				Freshness:       "2026-05",
				EvidenceGrade:   "survey",
				Confidence:      "medium",
			},
		},
	}

	runs, domainScores, observations := History(seed)

	if len(runs) != 2 || runs[1].RunID != "2026-q2" || runs[1].CompositeScore != 51.2 {
		t.Fatalf("unexpected explicit runs: %+v", runs)
	}
	if runs[0].IsPublicSeries || !runs[1].IsPublicSeries {
		t.Fatalf("unexpected public series flags: %+v", runs)
	}
	if len(domainScores) != 1 || domainScores[0].Score != 58.4 {
		t.Fatalf("unexpected explicit domain scores: %+v", domainScores)
	}
	if len(observations) != 1 || observations[0].SourceURL == "" || !observations[0].IncludedInScore {
		t.Fatalf("unexpected explicit observations: %+v", observations)
	}
}

func boolPtr(v bool) *bool {
	return &v
}
