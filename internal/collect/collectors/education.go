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
		{IndicatorName: "Students Using AI for Schoolwork", DomainID: "education", SourceName: "Pew Research Center", Err: fmt.Errorf("manual source: update overrides.yaml from pewresearch.org")},
		{IndicatorName: "AI-Graded Assessments", DomainID: "education", SourceName: "Gradescope/EdTech Reports", Err: fmt.Errorf("manual source: update overrides.yaml from Gradescope/edtech automated grading adoption data")},
		{IndicatorName: "Teachers Using AI for Work", DomainID: "education", SourceName: "OECD Digital Education Outlook", Err: fmt.Errorf("manual source: update overrides.yaml from OECD digital education reports")},
		{IndicatorName: "Student Papers 80%+ AI-Written", DomainID: "education", SourceName: "Turnitin AI Writing", Err: fmt.Errorf("manual source: update overrides.yaml from Turnitin AI writing reports")},
	}, nil
}
