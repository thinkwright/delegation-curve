package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

type LegalAICollector struct{}

func NewLegalAICollector() *LegalAICollector { return &LegalAICollector{} }

func (c *LegalAICollector) Name() string     { return "Legal AI Sources" }
func (c *LegalAICollector) DomainID() string { return "legal-ai" }

func (c *LegalAICollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{
		{IndicatorName: "AI Tool Adoption (BigLaw)", DomainID: "legal-ai", SourceName: "ALM Survey", Err: fmt.Errorf("manual source: update overrides.yaml from law.com/americanlawyer surveys")},
		{IndicatorName: "AI Tool Adoption (Solo/Small)", DomainID: "legal-ai", SourceName: "Clio Legal Trends", Err: fmt.Errorf("manual source: update overrides.yaml from clio.com/resources/legal-trends")},
		{IndicatorName: "AI-Assisted Document Review", DomainID: "legal-ai", SourceName: "EDRM/Relativity Survey", Err: fmt.Errorf("manual source: update overrides.yaml from EDRM/Relativity TAR adoption surveys")},
	}, nil
}
