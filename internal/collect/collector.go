package collect

import (
	"context"
	"time"

	"github.com/thinkwright/delegation-curve/internal/ingest"
)

// Collector fetches sub-indicator values from one external data source.
//
// Error convention: Return (results, nil) with Err set on individual
// CollectResult entries for per-indicator failures (e.g., manual stubs,
// partial fetch failures). Return (nil, err) only for infrastructure-level
// failures that prevent any collection attempt (e.g., context canceled).
type Collector interface {
	Name() string
	DomainID() string
	Collect(ctx context.Context) ([]CollectResult, error)
}

// CollectResult is what every collector returns per sub-indicator.
type CollectResult struct {
	IndicatorName string
	DomainID      string
	RawValue      float64
	Unit          string // "%" or "M"
	SourceName    string
	Freshness     string
	CollectedAt   time.Time
	Meta          ingest.SourceMeta
	Err           error // non-nil = fallback to override/cached
}
