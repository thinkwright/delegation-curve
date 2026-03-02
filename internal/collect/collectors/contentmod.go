package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

type ContentModCollector struct{}

func NewContentModCollector() *ContentModCollector { return &ContentModCollector{} }

func (c *ContentModCollector) Name() string     { return "Content Moderation Transparency Reports" }
func (c *ContentModCollector) DomainID() string { return "content-mod" }

func (c *ContentModCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{
		{IndicatorName: "Meta Automated Detection", DomainID: "content-mod", SourceName: "Meta Transparency Report", Err: fmt.Errorf("manual source: update overrides.yaml from transparency.meta.com")},
		{IndicatorName: "Google Automated Removal", DomainID: "content-mod", SourceName: "Google Transparency Report", Err: fmt.Errorf("manual source: update overrides.yaml from transparencyreport.google.com")},
		{IndicatorName: "TikTok Automated Detection", DomainID: "content-mod", SourceName: "TikTok Transparency Report", Err: fmt.Errorf("manual source: update overrides.yaml from tiktok.com/transparency")},
		{IndicatorName: "X/Twitter Automated Action", DomainID: "content-mod", SourceName: "X Transparency Report", Err: fmt.Errorf("manual source: update overrides.yaml from transparency.x.com")},
	}, nil
}
