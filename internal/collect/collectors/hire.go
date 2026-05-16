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
		{IndicatorName: "Orgs Using AI in Talent Acquisition", DomainID: "hire", SourceName: "ICIMS/Aptitude", Err: fmt.Errorf("manual source: update overrides.yaml from ICIMS/Aptitude AI adoption report")},
		{IndicatorName: "AI Screening Use Case Adoption", DomainID: "hire", SourceName: "ICIMS/Aptitude", Err: fmt.Errorf("manual source: update overrides.yaml from ICIMS/Aptitude AI adoption report")},
		{IndicatorName: "Broad AI Across Hiring Processes", DomainID: "hire", SourceName: "ICIMS/Aptitude", Err: fmt.Errorf("manual source: update overrides.yaml from ICIMS/Aptitude AI adoption report")},
		{IndicatorName: "AI Assessment Platform Reach", DomainID: "hire", SourceName: "Platform Aggregation", Err: fmt.Errorf("manual source: update overrides.yaml from aggregated platform disclosures (HireVue, Pymetrics, Codility, HackerRank)")},
	}, nil
}
