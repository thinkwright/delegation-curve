package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

type EducationCollector struct{}

func NewEducationCollector() *EducationCollector { return &EducationCollector{} }

func (c *EducationCollector) Name() string     { return "Education & Assessment Sources" }
func (c *EducationCollector) DomainID() string { return "education" }

func (c *EducationCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{
		{IndicatorName: "Students Using AI Tutors", DomainID: "education", SourceName: "Educause Survey", Err: fmt.Errorf("manual source: update overrides.yaml from educause.edu/research-and-publications")},
		{IndicatorName: "AI-Graded Assessments", DomainID: "education", SourceName: "Gradescope/EdTech Reports", Err: fmt.Errorf("manual source: update overrides.yaml from Gradescope/edtech automated grading adoption data")},
		{IndicatorName: "Faculty Using AI in Teaching", DomainID: "education", SourceName: "AAUP/Educause Survey", Err: fmt.Errorf("manual source: update overrides.yaml from AAUP/Educause faculty AI adoption surveys")},
	}, nil
}
