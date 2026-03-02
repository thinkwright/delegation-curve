package collect

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// OverrideFile is the top-level structure of overrides.yaml.
// Keys are domain IDs.
type OverrideFile map[string][]OverrideEntry

// OverrideEntry holds a manually-entered sub-indicator value.
type OverrideEntry struct {
	Name      string  `yaml:"name"`
	Value     float64 `yaml:"value"`
	Unit      string  `yaml:"unit"`
	Source    string  `yaml:"source"`
	Freshness string  `yaml:"freshness"`
	EnteredAt string  `yaml:"entered_at"`
	Note      string  `yaml:"note"`
}

// LoadOverrides reads and parses the overrides YAML file.
// Returns an empty map if the file does not exist.
func LoadOverrides(path string) (OverrideFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return OverrideFile{}, nil
		}
		return nil, fmt.Errorf("read overrides: %w", err)
	}
	var f OverrideFile
	if err := yaml.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("parse overrides: %w", err)
	}
	return f, nil
}
