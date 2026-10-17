package collect

import (
	"regexp"
	"testing"
)

// isoDate accepts the precisions sources disclose: YYYY, YYYY-MM, YYYY-MM-DD.
var isoDate = regexp.MustCompile(`^\d{4}(-\d{2}(-\d{2})?)?$`)

var (
	validGrades      = map[string]bool{"A": true, "B": true, "C": true, "D": true}
	validConfidences = map[string]bool{
		"low": true, "low-medium": true, "medium-low": true, "medium": true,
		"medium-high": true, "high": true,
	}
)

// TestCommittedOverridesCarryAuditTrail lints seed/overrides.yaml: every
// value that feeds the score must say where it came from, when it was
// measured, and how strong the evidence is.
func TestCommittedOverridesCarryAuditTrail(t *testing.T) {
	overrides, err := LoadOverrides("../../seed/overrides.yaml")
	if err != nil {
		t.Fatal(err)
	}
	if len(overrides) == 0 {
		t.Fatal("no overrides loaded")
	}

	configured := map[string]map[string]bool{}
	for _, d := range AllDomainConfigs() {
		configured[d.DomainID] = map[string]bool{}
		for _, ind := range d.Indicators {
			configured[d.DomainID][ind.Name] = true
		}
	}

	for domainID, entries := range overrides {
		for _, e := range entries {
			name := domainID + " / " + e.Name
			if !configured[domainID][e.Name] {
				t.Errorf("%s: not a configured indicator", name)
			}
			if !validGrades[e.EvidenceGrade] {
				t.Errorf("%s: evidence_grade %q, want A-D", name, e.EvidenceGrade)
			}
			if !validConfidences[e.Confidence] {
				t.Errorf("%s: confidence %q not recognized", name, e.Confidence)
			}
			if e.SourceURL == "" && e.Confidence != "low" {
				t.Errorf("%s: source_url is required unless the value is a low-confidence hold", name)
			}
			if e.SourceURL == "" && e.Transform == "" {
				t.Errorf("%s: held value without source_url must explain itself in transform", name)
			}
			for field, v := range map[string]string{
				"published_at": e.PublishedAt, "measurement_start": e.MeasurementStart,
				"measurement_end": e.MeasurementEnd, "entered_at": e.EnteredAt,
			} {
				if v != "" && !isoDate.MatchString(v) {
					t.Errorf("%s: %s %q is not an ISO date", name, field, v)
				}
			}
			if e.MeasurementEnd == "" && e.SourceURL != "" {
				t.Errorf("%s: measurement_end is required for sourced values", name)
			}
			if e.MeasurementStart != "" && e.MeasurementEnd != "" && e.MeasurementStart > e.MeasurementEnd {
				t.Errorf("%s: measurement_start %s after measurement_end %s", name, e.MeasurementStart, e.MeasurementEnd)
			}
		}
	}
}
