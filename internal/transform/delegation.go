package transform

import (
	"encoding/json"

	"github.com/thinkwright/delegation-curve/internal/ingest"
	"github.com/thinkwright/delegation-curve/internal/schema"
)

func Delegation(domains []ingest.DomainJSON) (
	[]schema.DelegationRow,
	[]schema.SubIndicatorRow,
	[]schema.DataSourceRow,
) {
	var dRows []schema.DelegationRow
	var sRows []schema.SubIndicatorRow
	var dsRows []schema.DataSourceRow

	for _, d := range domains {
		trendJSON, _ := json.Marshal(d.Trend)

		dRows = append(dRows, schema.DelegationRow{
			ID:            d.ID,
			Name:          d.Name,
			FullName:      d.FullName,
			Score:         d.Score,
			PreviousScore: d.PreviousScore,
			Trend:         string(trendJSON),
			Status:        d.Status,
			Weight:        d.Weight,
			Tier:          int32(d.Tier),
			Description:   d.Description,
		})

		for _, si := range d.SubIndicators {
			sRows = append(sRows, schema.SubIndicatorRow{
				DomainID:  d.ID,
				Name:      si.Name,
				Value:     si.Value,
				Unit:      si.Unit,
				Source:    si.Source,
				Freshness: si.Freshness,
			})
		}

		for _, ds := range d.DataSources {
			dsRows = append(dsRows, schema.DataSourceRow{
				DomainID: d.ID,
				Name:     ds.Name,
				Cadence:  ds.Cadence,
				Type:     ds.Type,
			})
		}
	}

	return dRows, sRows, dsRows
}
