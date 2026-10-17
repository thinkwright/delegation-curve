package collect

import (
	"fmt"
	"math"
	"os"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

// ScoreChange is one value that differs between two seeds.
type ScoreChange struct {
	DomainID  string // empty for the composite
	Indicator string // empty for domain and composite scores
	Before    float64
	After     float64
	Unit      string
	Missing   bool // indicator absent from the before seed
}

func (c ScoreChange) String() string {
	label := "composite"
	switch {
	case c.Indicator != "":
		label = c.DomainID + " / " + c.Indicator
	case c.DomainID != "":
		label = c.DomainID + " score"
	}
	if c.Missing {
		return fmt.Sprintf("%s: (new) → %g %s", label, c.After, c.Unit)
	}
	return fmt.Sprintf("%s: %g → %g %s (Δ%+.1f)", label, c.Before, c.After, c.Unit, c.After-c.Before)
}

// DiffScores lists composite, domain, and indicator values that differ
// between before and after. Indicator values are compared at the precision
// they are entered with; scores at one decimal.
func DiffScores(before, after *ingest.Seed) []ScoreChange {
	var changes []ScoreChange
	if math.Abs(before.Delegation.Composite.Current-after.Delegation.Composite.Current) >= 0.05 {
		changes = append(changes, ScoreChange{
			Before: before.Delegation.Composite.Current,
			After:  after.Delegation.Composite.Current,
		})
	}

	prevDomains := make(map[string]ingest.DomainJSON, len(before.Delegation.Domains))
	for _, d := range before.Delegation.Domains {
		prevDomains[d.ID] = d
	}
	for _, d := range after.Delegation.Domains {
		prev := prevDomains[d.ID]
		if math.Abs(prev.Score-d.Score) >= 0.05 {
			changes = append(changes, ScoreChange{DomainID: d.ID, Before: prev.Score, After: d.Score})
		}
		prevValues := make(map[string]float64, len(prev.SubIndicators))
		for _, si := range prev.SubIndicators {
			prevValues[si.Name] = si.Value
		}
		for _, si := range d.SubIndicators {
			v, ok := prevValues[si.Name]
			if !ok || v != si.Value {
				changes = append(changes, ScoreChange{
					DomainID: d.ID, Indicator: si.Name,
					Before: v, After: si.Value, Unit: si.Unit, Missing: !ok,
				})
			}
		}
	}
	return changes
}

func reportChanges(changes []ScoreChange) {
	if len(changes) == 0 {
		fmt.Fprintln(os.Stderr, "\n  No value or score changes.")
		return
	}
	fmt.Fprintf(os.Stderr, "\n  %d change(s):\n", len(changes))
	for _, c := range changes {
		fmt.Fprintf(os.Stderr, "    %s\n", c)
	}
}
