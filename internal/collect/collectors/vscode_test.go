package collectors

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// fakeMarketplace answers extension queries with fixed install counts;
// extensions missing from installs return no results.
func fakeMarketplace(t *testing.T, installs map[string]float64) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		body, _ := io.ReadAll(r.Body)
		var q struct {
			Filters []struct {
				Criteria []struct {
					Value string `json:"value"`
				} `json:"criteria"`
			} `json:"filters"`
		}
		if err := json.Unmarshal(body, &q); err != nil {
			t.Errorf("bad query: %v", err)
		}
		ext := q.Filters[0].Criteria[0].Value
		n, ok := installs[ext]
		if !ok {
			fmt.Fprint(w, `{"results":[{"extensions":[]}]}`)
			return
		}
		fmt.Fprintf(w, `{"results":[{"extensions":[{"statistics":[{"statisticName":"install","value":%v}]}]}]}`, n)
	}))
}

func TestVSCodeCollectorSumsAllExtensions(t *testing.T) {
	srv := fakeMarketplace(t, map[string]float64{"a.one": 2_000_000, "b.two": 500_000})
	defer srv.Close()
	c := &VSCodeCollector{client: srv.Client(), extensions: []string{"a.one", "b.two"}, endpoint: srv.URL}

	results, err := c.Collect(context.Background())
	if err != nil || len(results) != 1 {
		t.Fatalf("Collect = %v, %v", results, err)
	}
	r := results[0]
	if r.Err != nil || r.RawValue != 2.5 || r.Unit != "M" {
		t.Fatalf("result = %+v", r)
	}
	if r.Meta.EvidenceGrade != "D" || r.Meta.SourceURL == "" || r.Meta.MeasurementEnd == "" {
		t.Fatalf("missing audit metadata: %+v", r.Meta)
	}
}

func TestVSCodeCollectorRejectsPartialSum(t *testing.T) {
	srv := fakeMarketplace(t, map[string]float64{"a.one": 2_000_000})
	defer srv.Close()
	c := &VSCodeCollector{client: srv.Client(), extensions: []string{"a.one", "gone.ext"}, endpoint: srv.URL}

	results, err := c.Collect(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if results[0].Err == nil || !strings.Contains(results[0].Err.Error(), "1 of 2 extensions failed") {
		t.Fatalf("partial result should fall back, got %+v", results[0])
	}
}
