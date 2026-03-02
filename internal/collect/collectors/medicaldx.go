package collectors

import (
	"context"
	"fmt"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

type MedicalDxCollector struct{}

func NewMedicalDxCollector() *MedicalDxCollector { return &MedicalDxCollector{} }

func (c *MedicalDxCollector) Name() string     { return "Medical Diagnosis Sources" }
func (c *MedicalDxCollector) DomainID() string { return "medical-dx" }

func (c *MedicalDxCollector) Collect(_ context.Context) ([]collect.CollectResult, error) {
	return []collect.CollectResult{
		{IndicatorName: "FDA-Cleared Diagnostic AI Devices", DomainID: "medical-dx", SourceName: "FDA AI/ML Database", Err: fmt.Errorf("manual source: update overrides.yaml from fda.gov AI/ML database (filter to radiology/pathology/cardiology)")},
		{IndicatorName: "Radiology AI Adoption", DomainID: "medical-dx", SourceName: "RSNA Survey", Err: fmt.Errorf("manual source: update overrides.yaml from rsna.org AI survey")},
		{IndicatorName: "AI-Assisted Diagnosis Rate", DomainID: "medical-dx", SourceName: "AMA Survey", Err: fmt.Errorf("manual source: update overrides.yaml from ama-assn.org physician AI sentiment report (diagnosis-specific breakout)")},
		{IndicatorName: "Pathology AI Adoption", DomainID: "medical-dx", SourceName: "CAP Survey", Err: fmt.Errorf("manual source: update overrides.yaml from cap.org digital pathology surveys")},
	}, nil
}
