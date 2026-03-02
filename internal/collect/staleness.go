package collect

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// CollectLog records the result of a collection run.
type CollectLog struct {
	RunAt   time.Time              `json:"run_at"`
	Entries map[string][]LogEntry `json:"entries"` // keyed by domain ID
}

// LogEntry records a single sub-indicator's collection outcome.
type LogEntry struct {
	IndicatorName string    `json:"indicator_name"`
	SourceName    string    `json:"source_name"`
	Value         float64   `json:"value"`
	Unit          string    `json:"unit"`
	Freshness     string    `json:"freshness"`
	CollectedAt   time.Time `json:"collected_at"`
	Method        string    `json:"method"` // "auto", "override", "cached", "missing"
	Error         string    `json:"error,omitempty"`
	SourceURL     string    `json:"source_url,omitempty"`
}

// DefaultStalenessThreshold is the fallback when no cadence metadata is available.
const DefaultStalenessThreshold = 400 * 24 * time.Hour // ~13 months (annual cadence)

// LoadCollectLog reads the collect log from disk.
// Returns an empty log if the file does not exist.
func LoadCollectLog(path string) (*CollectLog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &CollectLog{Entries: map[string][]LogEntry{}}, nil
		}
		return nil, fmt.Errorf("read collect log: %w", err)
	}
	var log CollectLog
	if err := json.Unmarshal(data, &log); err != nil {
		return nil, fmt.Errorf("parse collect log: %w", err)
	}
	return &log, nil
}

// Save atomically writes the collect log to disk.
func (l *CollectLog) Save(path string) error {
	data, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal collect log: %w", err)
	}
	return atomicWrite(path, data)
}

// CheckOverrideStaleness returns warnings for any override entries whose
// freshness date exceeds the cadence-aware staleness threshold for that indicator.
// It builds a lookup from AllDomainConfigs to resolve each indicator's native cadence.
func CheckOverrideStaleness(overrides OverrideFile) []string {
	// Build indicator name → cadence lookup from all domain configs.
	cadenceOf := make(map[string]Cadence)
	for _, dc := range AllDomainConfigs() {
		for _, ind := range dc.Indicators {
			cadenceOf[ind.Name] = ind.Cadence
		}
	}

	var warnings []string
	now := time.Now()

	for domainID, entries := range overrides {
		for _, e := range entries {
			age := overrideAge(e, now)
			threshold := DefaultStalenessThreshold
			if c, ok := cadenceOf[e.Name]; ok {
				threshold = c.StalenessThreshold()
			}
			if age > threshold {
				warnings = append(warnings, fmt.Sprintf(
					"STALE: %s/%s override is %.0f days old (freshness: %q, entered: %q)",
					domainID, e.Name, age.Hours()/24, e.Freshness, e.EnteredAt,
				))
			}
		}
	}
	return warnings
}

// overrideAge computes how old an override entry is, trying EnteredAt first,
// then Freshness. Returns zero duration if neither is parseable.
func overrideAge(e OverrideEntry, now time.Time) time.Duration {
	for _, dateStr := range []string{e.EnteredAt, e.Freshness} {
		if dateStr == "" {
			continue
		}
		for _, layout := range []string{"2006-01-02", "2006", "Q1 2006", "Q2 2006", "Q3 2006", "Q4 2006"} {
			if t, err := time.Parse(layout, dateStr); err == nil {
				return now.Sub(t)
			}
		}
	}
	return 0
}

func atomicWrite(path string, data []byte) error {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".tmp-*")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpName := tmp.Name()

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return fmt.Errorf("write temp file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return fmt.Errorf("sync temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("close temp file: %w", err)
	}
	if err := os.Rename(tmpName, path); err != nil {
		os.Remove(tmpName)
		return fmt.Errorf("rename temp file: %w", err)
	}
	return nil
}
