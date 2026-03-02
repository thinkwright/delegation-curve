package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

type CreditCollector struct{}

func NewCreditCollector() *CreditCollector { return &CreditCollector{} }

func (c *CreditCollector) Name() string     { return "Credit & Lending Sources" }
func (c *CreditCollector) DomainID() string { return "credit" }

func (c *CreditCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{
		{IndicatorName: "AI-Underwritten Loan Volume", DomainID: "credit", SourceName: "Fintech Filings", Err: fmt.Errorf("manual source: update overrides.yaml from fintech company filings")},
		{IndicatorName: "Fintech Lending Market Share", DomainID: "credit", SourceName: "Industry Reports", Err: fmt.Errorf("manual source: update overrides.yaml from fintech market share reports (Upstart, SoFi, LendingClub filings)")},
		{IndicatorName: "AI Credit Decisioning (Banks)", DomainID: "credit", SourceName: "OCC Survey", Err: fmt.Errorf("manual source: update overrides.yaml from occ.gov surveys (credit-specific ML adoption)")},
	}, nil
}
