package transform

import (
	"fmt"
	"sort"

	"github.com/thinkwright/delegation-curve/internal/collect"
	"github.com/thinkwright/delegation-curve/internal/ingest"
	"github.com/thinkwright/delegation-curve/internal/schema"
)

const legacyMethodologyVersion = "legacy-trend-v1"

func History(seed *ingest.Seed) (
	[]schema.AnalysisRunRow,
	[]schema.DomainScoreRow,
	[]schema.IndicatorObservationRow,
) {
	if len(seed.AnalysisRuns) > 0 {
		return explicitHistory(seed)
	}
	return legacyHistory(seed)
}

func explicitHistory(seed *ingest.Seed) (
	[]schema.AnalysisRunRow,
	[]schema.DomainScoreRow,
	[]schema.IndicatorObservationRow,
) {
	runRows := make([]schema.AnalysisRunRow, 0, len(seed.AnalysisRuns))
	for _, r := range seed.AnalysisRuns {
		isPublicSeries := true
		if r.IsPublicSeries != nil {
			isPublicSeries = *r.IsPublicSeries
		}
		runRows = append(runRows, schema.AnalysisRunRow{
			RunID:              r.RunID,
			Label:              r.Label,
			PublishedAt:        r.PublishedAt,
			MeasurementPeriod:  r.MeasurementPeriod,
			MeasurementYear:    int32(r.MeasurementYear),
			MethodologyVersion: r.MethodologyVersion,
			CompositeScore:     r.CompositeScore,
			Notes:              r.Notes,
			IsCurrent:          r.IsCurrent,
			IsPublicSeries:     isPublicSeries,
		})
	}

	domainRows := make([]schema.DomainScoreRow, 0, len(seed.DomainScores))
	for _, d := range seed.DomainScores {
		domainRows = append(domainRows, schema.DomainScoreRow{
			RunID:    d.RunID,
			DomainID: d.DomainID,
			Score:    d.Score,
			Weight:   d.Weight,
			Status:   d.Status,
		})
	}

	indicatorRows := make([]schema.IndicatorObservationRow, 0, len(seed.IndicatorObservations))
	for _, o := range seed.IndicatorObservations {
		indicatorRows = append(indicatorRows, schema.IndicatorObservationRow{
			RunID:           o.RunID,
			DomainID:        o.DomainID,
			IndicatorName:   o.IndicatorName,
			RawValue:        o.RawValue,
			NormalizedValue: o.NormalizedValue,
			Unit:            o.Unit,
			Weight:          o.Weight,
			IncludedInScore: o.IncludedInScore,
			Source:          o.Source,
			SourceURL:       o.SourceURL,
			Freshness:       o.Freshness,
			EvidenceGrade:   o.EvidenceGrade,
			Confidence:      o.Confidence,
		})
	}

	sortAnalysisRuns(runRows)
	sortDomainScores(domainRows)
	sortIndicatorObservations(indicatorRows)

	return runRows, domainRows, indicatorRows
}

func legacyHistory(seed *ingest.Seed) (
	[]schema.AnalysisRunRow,
	[]schema.DomainScoreRow,
	[]schema.IndicatorObservationRow,
) {
	composite := seed.Delegation.Composite
	startYear := composite.DataYear - len(composite.Trend) + 1

	runRows := make([]schema.AnalysisRunRow, 0, len(composite.Trend))
	for i, score := range composite.Trend {
		year := startYear + i
		runRows = append(runRows, schema.AnalysisRunRow{
			RunID:              legacyRunID(year),
			Label:              fmt.Sprintf("%d baseline", year),
			PublishedAt:        publishedAtForLegacyRun(seed, year),
			MeasurementPeriod:  fmt.Sprintf("%d", year),
			MeasurementYear:    int32(year),
			MethodologyVersion: legacyMethodologyVersion,
			CompositeScore:     score,
			Notes:              "Derived from legacy trend arrays in seed.json.",
			IsCurrent:          year == composite.DataYear,
			IsPublicSeries:     true,
		})
	}

	var domainRows []schema.DomainScoreRow
	for _, d := range seed.Delegation.Domains {
		domainStartYear := composite.DataYear - len(d.Trend) + 1
		for i, score := range d.Trend {
			year := domainStartYear + i
			domainRows = append(domainRows, schema.DomainScoreRow{
				RunID:    legacyRunID(year),
				DomainID: d.ID,
				Score:    score,
				Weight:   d.Weight,
				Status:   collect.ClassifyStatus(score),
			})
		}
	}

	indicatorRows := CurrentIndicatorObservations(seed, legacyRunID(composite.DataYear))
	sortAnalysisRuns(runRows)
	sortDomainScores(domainRows)
	sortIndicatorObservations(indicatorRows)
	return runRows, domainRows, indicatorRows
}

func CurrentIndicatorObservations(seed *ingest.Seed, runID string) []schema.IndicatorObservationRow {
	configs := map[string]collect.DomainConfig{}
	for _, cfg := range collect.AllDomainConfigs() {
		configs[cfg.DomainID] = cfg
	}

	var rows []schema.IndicatorObservationRow
	for _, d := range seed.Delegation.Domains {
		indicatorConfig := map[string]collect.IndicatorConfig{}
		if cfg, ok := configs[d.ID]; ok {
			for _, ind := range cfg.Indicators {
				indicatorConfig[ind.Name] = ind
			}
		}

		for _, si := range d.SubIndicators {
			indCfg, included := indicatorConfig[si.Name]
			normalized := 0.0
			weight := 0.0
			if included {
				normalized = collect.Normalize(si.Value, indCfg.NormConfig)
				weight = indCfg.Weight
			}

			rows = append(rows, schema.IndicatorObservationRow{
				RunID:           runID,
				DomainID:        d.ID,
				IndicatorName:   si.Name,
				RawValue:        si.Value,
				NormalizedValue: normalized,
				Unit:            si.Unit,
				Weight:          weight,
				IncludedInScore: included,
				Source:          si.Source,
				Freshness:       si.Freshness,
			})
		}
	}
	return rows
}

func legacyRunID(year int) string {
	return fmt.Sprintf("legacy-%d", year)
}

func publishedAtForLegacyRun(seed *ingest.Seed, year int) string {
	if year == seed.Delegation.Composite.DataYear {
		return seed.Delegation.Composite.LastUpdated
	}
	return ""
}

func sortAnalysisRuns(rows []schema.AnalysisRunRow) {
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].MeasurementYear != rows[j].MeasurementYear {
			return rows[i].MeasurementYear < rows[j].MeasurementYear
		}
		if rows[i].PublishedAt != rows[j].PublishedAt {
			return rows[i].PublishedAt < rows[j].PublishedAt
		}
		return rows[i].RunID < rows[j].RunID
	})
}

func sortDomainScores(rows []schema.DomainScoreRow) {
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].RunID != rows[j].RunID {
			return rows[i].RunID < rows[j].RunID
		}
		return rows[i].DomainID < rows[j].DomainID
	})
}

func sortIndicatorObservations(rows []schema.IndicatorObservationRow) {
	sort.SliceStable(rows, func(i, j int) bool {
		if rows[i].RunID != rows[j].RunID {
			return rows[i].RunID < rows[j].RunID
		}
		if rows[i].DomainID != rows[j].DomainID {
			return rows[i].DomainID < rows[j].DomainID
		}
		return rows[i].IndicatorName < rows[j].IndicatorName
	})
}
