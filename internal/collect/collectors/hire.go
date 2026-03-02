package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

type HireCollector struct{}

func NewHireCollector() *HireCollector { return &HireCollector{} }

func (c *HireCollector) Name() string     { return "Recruitment & Screening Sources" }
func (c *HireCollector) DomainID() string { return "hire" }

func (c *HireCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{
		{IndicatorName: "Orgs Using AI Screening", DomainID: "hire", SourceName: "SHRM Survey", Err: fmt.Errorf("manual source: update overrides.yaml from shrm.org research")},
		{IndicatorName: "AI-Screened Applications", DomainID: "hire", SourceName: "SHRM/LinkedIn Data", Err: fmt.Errorf("manual source: update overrides.yaml from SHRM/LinkedIn hiring AI adoption data")},
		{IndicatorName: "AI Assessment Platform Reach", DomainID: "hire", SourceName: "Platform Aggregation", Err: fmt.Errorf("manual source: update overrides.yaml from aggregated platform disclosures (HireVue, Pymetrics, Codility, HackerRank)")},
	}, nil
}
