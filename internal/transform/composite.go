package transform

import (
	"encoding/json"

	"github.com/thinkwright/delegation-curve/internal/ingest"
	"github.com/thinkwright/delegation-curve/internal/schema"
)

func Meta(seed *ingest.Seed) schema.MetaRow {
	trendJSON, _ := json.Marshal(seed.Delegation.Composite.Trend)

	// Find highest-scoring domain
	var highName string
	var highScore float64
	for _, d := range seed.Delegation.Domains {
		if d.Score > highScore {
			highScore = d.Score
			highName = d.Name
		}
	}

	return schema.MetaRow{
		DelegationComposite: seed.Delegation.Composite.Current,
		DelegationPrevious:  seed.Delegation.Composite.Previous,
		DelegationDelta:     seed.Delegation.Composite.Delta,
		DelegationTrend:     string(trendJSON),
		DomainsTracked:      int32(len(seed.Delegation.Domains)),
		HighestDomainName:   highName,
		HighestDomainScore:  highScore,
		DataFreshness:       "Q4 2025",
		LastUpdated:         seed.Delegation.Composite.LastUpdated,
		DataYear:            int32(seed.Delegation.Composite.DataYear),
	}
}
