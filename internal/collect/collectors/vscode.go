package collectors

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/thinkwright/delegation-curve/internal/collect"
)

// aiExtensions lists the major AI coding assistant extensions to aggregate.
var aiExtensions = []string{
	"GitHub.copilot",
	"Codeium.codeium",
	"TabNine.tabnine-vscode",
	"Continue.continue",
	"Supermaven.supermaven",
}

// VSCodeCollector fetches aggregate install counts for AI coding extensions
// from the VS Code Marketplace API.
type VSCodeCollector struct {
	client     *http.Client
	extensions []string
}

func NewVSCodeCollector() *VSCodeCollector {
	return &VSCodeCollector{
		client:     &http.Client{Timeout: 15 * time.Second},
		extensions: aiExtensions,
	}
}

func (c *VSCodeCollector) Name() string     { return "VS Code Marketplace" }
func (c *VSCodeCollector) DomainID() string { return "code-gen" }

func (c *VSCodeCollector) Collect(ctx context.Context) ([]collect.CollectResult, error) {
	var totalInstalls int64
	var fetchErrors []string

	for _, ext := range c.extensions {
		installs, err := c.fetchInstallCount(ctx, ext)
		if err != nil {
			fetchErrors = append(fetchErrors, fmt.Sprintf("%s: %v", ext, err))
			continue
		}
		totalInstalls += installs
	}

	// Fail only if we got zero extensions at all.
	if totalInstalls == 0 && len(fetchErrors) == len(c.extensions) {
		return []collect.CollectResult{{
			IndicatorName: "IDE AI Extension Installs",
			DomainID:      "code-gen",
			SourceName:    "VS Code Marketplace",
			Err:           fmt.Errorf("all extensions failed: %v", fetchErrors),
		}}, nil
	}

	millions := float64(totalInstalls) / 1_000_000
	return []collect.CollectResult{{
		IndicatorName: "IDE AI Extension Installs",
		DomainID:      "code-gen",
		RawValue:      millions,
		Unit:          "M",
		SourceName:    "VS Code Marketplace",
		Freshness:     time.Now().Format("2006-01-02"),
		SourceURL:     "https://marketplace.visualstudio.com/",
	}}, nil
}

func (c *VSCodeCollector) fetchInstallCount(ctx context.Context, extension string) (int64, error) {
	body := map[string]any{
		"filters": []map[string]any{{
			"criteria": []map[string]any{{
				"filterType": 7,
				"value":      extension,
			}},
			"pageNumber": 1,
			"pageSize":   1,
		}},
		"flags": 914,
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST",
		"https://marketplace.visualstudio.com/_apis/public/gallery/extensionquery",
		bytes.NewReader(payload))
	if err != nil {
		return 0, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json;api-version=3.0-preview.1")

	resp, err := c.client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("marketplace API returned %d for %s", resp.StatusCode, extension)
	}

	data, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return 0, err
	}

	var result struct {
		Results []struct {
			Extensions []struct {
				Statistics []struct {
					StatisticName string  `json:"statisticName"`
					Value         float64 `json:"value"`
				} `json:"statistics"`
			} `json:"extensions"`
		} `json:"results"`
	}
	if err := json.Unmarshal(data, &result); err != nil {
		return 0, fmt.Errorf("parse marketplace response: %w", err)
	}

	if len(result.Results) == 0 || len(result.Results[0].Extensions) == 0 {
		return 0, fmt.Errorf("no extension found for %s", extension)
	}

	for _, stat := range result.Results[0].Extensions[0].Statistics {
		if stat.StatisticName == "install" {
			if stat.Value <= 0 || stat.Value > 1e10 {
				return 0, fmt.Errorf("install count out of range for %s: %.0f", extension, stat.Value)
			}
			return int64(stat.Value), nil
		}
	}

	return 0, fmt.Errorf("install statistic not found for %s", extension)
}
