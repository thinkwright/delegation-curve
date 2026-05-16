package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

type SupportCollector struct{}

func NewSupportCollector() *SupportCollector { return &SupportCollector{} }

func (c *SupportCollector) Name() string     { return "Customer Support Surveys" }
func (c *SupportCollector) DomainID() string { return "support" }

func (c *SupportCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{
		{IndicatorName: "Cases Handled by AI", DomainID: "support", SourceName: "Salesforce State of Service", Err: fmt.Errorf("manual source: update overrides.yaml from salesforce.com/resources/research-reports/state-of-service")},
		{IndicatorName: "Bot Deflection Rate", DomainID: "support", SourceName: "Intercom Trends", Err: fmt.Errorf("manual source: update overrides.yaml from intercom.com/customer-support-trends")},
		{IndicatorName: "Production AI Customer Communications Agents", DomainID: "support", SourceName: "Sinch AI Production Paradox", Err: fmt.Errorf("manual source: update overrides.yaml from sinch.com/news")},
		{IndicatorName: "Mature AI Support Deployment", DomainID: "support", SourceName: "Intercom Customer Service Transformation", Err: fmt.Errorf("manual source: update overrides.yaml from intercom.com/customer-transformation-report")},
	}, nil
}
