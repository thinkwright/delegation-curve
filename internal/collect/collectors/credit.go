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
		{IndicatorName: "AI-Underwritten Personal Loan Proxy", DomainID: "credit", SourceName: "Upstart and TransUnion", Err: fmt.Errorf("manual source: update overrides.yaml from Upstart filings and TransUnion CIIR")},
		{IndicatorName: "AI Credit Decisioning (Banks)", DomainID: "credit", SourceName: "OCC Survey", Err: fmt.Errorf("manual source: update overrides.yaml from occ.gov surveys (credit-specific ML adoption)")},
	}, nil
}
