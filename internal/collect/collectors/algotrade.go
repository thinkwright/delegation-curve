package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

type AlgoTradeCollector struct{}

func NewAlgoTradeCollector() *AlgoTradeCollector { return &AlgoTradeCollector{} }

func (c *AlgoTradeCollector) Name() string     { return "Algorithmic Trading Surveys" }
func (c *AlgoTradeCollector) DomainID() string { return "algo-trade" }

func (c *AlgoTradeCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{
		{IndicatorName: "FX Electronic Trading Share", DomainID: "algo-trade", SourceName: "BIS Triennial Survey", Err: fmt.Errorf("manual source: update overrides.yaml from bis.org/statistics")},
		{IndicatorName: "Buy-Side AI Trade Execution Adoption", DomainID: "algo-trade", SourceName: "Coalition Greenwich", Err: fmt.Errorf("manual source: update overrides.yaml from greenwich.com")},
	}, nil
}
