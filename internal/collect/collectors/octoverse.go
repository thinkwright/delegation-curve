package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

// OctoverseCollector is a manual stub for the GitHub Octoverse annual report.
// The report is an HTML blog post with varying structure year-to-year,
// making scraping fragile. Values should be entered in overrides.yaml.
type OctoverseCollector struct{}

func NewOctoverseCollector() *OctoverseCollector { return &OctoverseCollector{} }

func (c *OctoverseCollector) Name() string     { return "GitHub Octoverse" }
func (c *OctoverseCollector) DomainID() string { return "code-gen" }

func (c *OctoverseCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{{
		IndicatorName: "Copilot Code Acceptance",
		DomainID:      "code-gen",
		SourceName:    "GitHub Octoverse",
		Err:           fmt.Errorf("manual source: update overrides.yaml from github.blog/news-insights/octoverse/"),
	}}, nil
}
