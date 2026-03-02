package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

// StackOverflowCollector is a manual stub for the Stack Overflow Developer Survey.
// The survey CSV is ~90MB with column names that change annually, making
// full automation fragile. Values should be entered in overrides.yaml.
type StackOverflowCollector struct{}

func NewStackOverflowCollector() *StackOverflowCollector { return &StackOverflowCollector{} }

func (c *StackOverflowCollector) Name() string     { return "Stack Overflow Survey" }
func (c *StackOverflowCollector) DomainID() string { return "code-gen" }

func (c *StackOverflowCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{{
		IndicatorName: "Developer AI Tool Usage",
		DomainID:      "code-gen",
		SourceName:    "Stack Overflow Survey",
		Err:           fmt.Errorf("manual source: update overrides.yaml from survey.stackoverflow.co"),
	}}, nil
}
