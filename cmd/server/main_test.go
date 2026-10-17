package main

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"testing/fstest"
)

func gz(t *testing.T, s string) []byte {
	t.Helper()
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	io.WriteString(w, s)
	w.Close()
	return b.Bytes()
}

func testSite(t *testing.T) http.Handler {
	t.Helper()
	fsys := fstest.MapFS{
		"index.html":                   {Data: []byte("home")},
		"about.html":                   {Data: []byte("about")},
		"delegation/code-gen.html":     {Data: []byte("code-gen")},
		"404.html":                     {Data: []byte("not found page")},
		"data/meta.json":               {Data: []byte(`{"a":1}`)},
		"data/meta.json.gz":            {Data: gz(t, `{"a":1}`)},
		"_app/immutable/entry/app.js":  {Data: []byte("js")},
		"data/delegation.parquet":      {Data: []byte("PAR1")},
		"delegation/code-gen/ignored/": {Mode: 0o755 | 1<<31}, // directory entry
	}
	return securityHeaders(siteHandler(http.FS(fsys)))
}

func get(t *testing.T, h http.Handler, target string, header ...string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, target, nil)
	for i := 0; i+1 < len(header); i += 2 {
		req.Header.Set(header[i], header[i+1])
	}
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, req)
	return rec
}

func TestSiteHandlerRoutes(t *testing.T) {
	h := testSite(t)
	tests := []struct {
		target string
		status int
		body   string
	}{
		{"/", 200, "home"},
		{"/about", 200, "about"},
		{"/delegation/code-gen", 200, "code-gen"},
		{"/delegation/unknown", 404, "not found page"},
		{"/.git/config", 404, "not found page"},
		{"/data/missing.parquet", 404, "not found page"},
	}
	for _, tt := range tests {
		rec := get(t, h, tt.target)
		if rec.Code != tt.status || rec.Body.String() != tt.body {
			t.Errorf("GET %s = %d %q, want %d %q", tt.target, rec.Code, rec.Body.String(), tt.status, tt.body)
		}
	}
}

func TestSiteHandlerRedirectsTrailingSlash(t *testing.T) {
	rec := get(t, testSite(t), "/about/?x=1")
	if rec.Code != http.StatusMovedPermanently || rec.Header().Get("Location") != "/about?x=1" {
		t.Fatalf("GET /about/ = %d Location %q", rec.Code, rec.Header().Get("Location"))
	}
}

func TestSiteHandlerServesPrecompressed(t *testing.T) {
	h := testSite(t)
	rec := get(t, h, "/data/meta.json", "Accept-Encoding", "gzip, br")
	if rec.Header().Get("Content-Encoding") != "gzip" || rec.Header().Get("Vary") != "Accept-Encoding" {
		t.Fatalf("gzip not negotiated: %v", rec.Header())
	}
	r, err := gzip.NewReader(rec.Body)
	if err != nil {
		t.Fatal(err)
	}
	if body, _ := io.ReadAll(r); string(body) != `{"a":1}` {
		t.Fatalf("decompressed body %q", body)
	}

	rec = get(t, h, "/data/meta.json", "Accept-Encoding", "gzip;q=0")
	if rec.Header().Get("Content-Encoding") != "" || rec.Body.String() != `{"a":1}` {
		t.Fatalf("q=0 should refuse gzip: %v %q", rec.Header(), rec.Body.String())
	}
}

func TestSiteHandlerHeaders(t *testing.T) {
	h := testSite(t)
	tests := []struct {
		target, cache, contentType string
	}{
		{"/", "no-cache", "text/html; charset=utf-8"},
		{"/_app/immutable/entry/app.js", "public, max-age=31536000, immutable", "application/javascript; charset=utf-8"},
		{"/data/delegation.parquet", "public, max-age=300, stale-while-revalidate=86400", "application/vnd.apache.parquet"},
		{"/nope", "no-cache", "text/html; charset=utf-8"},
	}
	for _, tt := range tests {
		rec := get(t, h, tt.target)
		if got := rec.Header().Get("Cache-Control"); got != tt.cache {
			t.Errorf("%s Cache-Control = %q, want %q", tt.target, got, tt.cache)
		}
		if got := rec.Header().Get("Content-Type"); got != tt.contentType {
			t.Errorf("%s Content-Type = %q, want %q", tt.target, got, tt.contentType)
		}
		if rec.Header().Get("X-Content-Type-Options") != "nosniff" {
			t.Errorf("%s missing security headers", tt.target)
		}
	}
}
