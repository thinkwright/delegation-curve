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
		{IndicatorName: "US Equities Algo Volume", DomainID: "algo-trade", SourceName: "SEC Market Structure", Err: fmt.Errorf("manual source: update overrides.yaml from sec.gov/market-structure")},
		{IndicatorName: "FX Algo Trading", DomainID: "algo-trade", SourceName: "BIS Triennial Survey", Err: fmt.Errorf("manual source: update overrides.yaml from bis.org/statistics")},
		{IndicatorName: "Options Algo Volume", DomainID: "algo-trade", SourceName: "CBOE Data", Err: fmt.Errorf("manual source: update overrides.yaml from cboe.com/data")},
		{IndicatorName: "Institutional AI Adoption", DomainID: "algo-trade", SourceName: "Greenwich Associates", Err: fmt.Errorf("manual source: update overrides.yaml from greenwich.com")},
	}, nil
}
