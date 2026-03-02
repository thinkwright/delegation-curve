package ingest

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type Seed struct {
	GeneratedAt string         `json:"generated_at"`
	Delegation  DelegationSeed `json:"delegation"`
}

type DelegationSeed struct {
	Composite CompositeDelegation `json:"composite"`
	Domains   []DomainJSON        `json:"domains"`
}

type CompositeDelegation struct {
	Current     float64   `json:"current"`
	Previous    float64   `json:"previous"`
	Delta       float64   `json:"delta"`
	Trend       []float64 `json:"trend"`
	LastUpdated string    `json:"last_updated"`
	DataYear    int       `json:"data_year"`
}

type DomainJSON struct {
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	FullName      string             `json:"full_name"`
	Score         float64            `json:"score"`
	PreviousScore float64            `json:"previous_score"`
	Trend         []float64          `json:"trend"`
	Status        string             `json:"status"`
	Weight        float64            `json:"weight"`
	Tier          int                `json:"tier"`
	Description   string             `json:"description"`
	SubIndicators []SubIndicatorJSON `json:"sub_indicators"`
	DataSources   []DataSourceJSON   `json:"data_sources"`
}

type SubIndicatorJSON struct {
	Name      string  `json:"name"`
	Value     float64 `json:"value"`
	Unit      string  `json:"unit"`
	Source    string  `json:"source"`
	Freshness string  `json:"freshness"`
}

type DataSourceJSON struct {
	Name    string `json:"name"`
	Cadence string `json:"cadence"`
	Type    string `json:"type"`
}

func ReadSeed(path string) (*Seed, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	var seed Seed
	if err := json.Unmarshal(data, &seed); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}
	return &seed, nil
}

// WriteSeed atomically writes the seed to disk via temp file + rename.
func WriteSeed(path string, seed *Seed) error {
	seed.GeneratedAt = time.Now().UTC().Format(time.RFC3339)
	data, err := json.MarshalIndent(seed, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal seed: %w", err)
	}
	return atomicWrite(path, data)
}

// atomicWrite writes data to a temp file then renames it into place.
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
