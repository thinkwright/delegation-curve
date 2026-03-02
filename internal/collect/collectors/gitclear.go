package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

// GitClearCollector is a manual stub for GitClear code quality analysis.
// GitClear's data requires a paid API subscription, so values
// must be entered manually in overrides.yaml.
type GitClearCollector struct{}

func NewGitClearCollector() *GitClearCollector { return &GitClearCollector{} }

func (c *GitClearCollector) Name() string     { return "GitClear Analysis" }
func (c *GitClearCollector) DomainID() string { return "code-gen" }

func (c *GitClearCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{{
		IndicatorName: "AI-Assisted Commits (OSS)",
		DomainID:      "code-gen",
		SourceName:    "GitClear Analysis",
		Err:           fmt.Errorf("manual source: update overrides.yaml from gitclear.com (requires API subscription)"),
	}}, nil
}
