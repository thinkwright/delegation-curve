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
		{IndicatorName: "AI Resolution Rate", DomainID: "support", SourceName: "Zendesk CX Trends", Err: fmt.Errorf("manual source: update overrides.yaml from zendesk.com/cx-trends")},
		{IndicatorName: "Bot Deflection Rate", DomainID: "support", SourceName: "Intercom Trends", Err: fmt.Errorf("manual source: update overrides.yaml from intercom.com/customer-support-trends")},
		{IndicatorName: "Orgs Using AI Support", DomainID: "support", SourceName: "Salesforce State of Service", Err: fmt.Errorf("manual source: update overrides.yaml from salesforce.com/resources/research-reports/state-of-service")},
		{IndicatorName: "AI Copilot Adoption (Agents)", DomainID: "support", SourceName: "Zendesk/Salesforce Reports", Err: fmt.Errorf("manual source: update overrides.yaml from Zendesk/Salesforce agent AI copilot adoption reports")},
	}, nil
}
