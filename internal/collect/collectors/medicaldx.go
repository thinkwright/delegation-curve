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
		{IndicatorName: "FDA AI-Enabled Medical Devices", DomainID: "medical-dx", SourceName: "FDA AI-Enabled Medical Devices", Err: fmt.Errorf("manual source: update overrides.yaml from fda.gov AI-enabled medical devices CSV")},
		{IndicatorName: "Radiology or Imaging AI Adoption", DomainID: "medical-dx", SourceName: "KLAS Global Imaging AI", Err: fmt.Errorf("manual source: update overrides.yaml from KLAS imaging AI reports")},
		{IndicatorName: "AI-Assisted Diagnosis Rate", DomainID: "medical-dx", SourceName: "AMA Survey", Err: fmt.Errorf("manual source: update overrides.yaml from ama-assn.org physician AI sentiment report (diagnosis-specific breakout)")},
		{IndicatorName: "Pathology AI Adoption", DomainID: "medical-dx", SourceName: "CAP Survey", Err: fmt.Errorf("manual source: update overrides.yaml from cap.org digital pathology surveys")},
	}, nil
}
