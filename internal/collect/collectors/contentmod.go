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
		{IndicatorName: "YouTube Automated Flagging", DomainID: "content-mod", SourceName: "YouTube Transparency Report", Err: fmt.Errorf("manual source: update overrides.yaml from transparencyreport.google.com")},
		{IndicatorName: "TikTok Automated Enforcement", DomainID: "content-mod", SourceName: "TikTok DSA Transparency Report", Err: fmt.Errorf("manual source: update overrides.yaml from tiktok.com/transparency or DSA report")},
	}, nil
}
