package main

import (
	"embed"
	"flag"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

//go:embed all:static
var staticFiles embed.FS

func main() {
	port := flag.String("port", getEnv("PORT", "8080"), "Port to listen on")
	flag.Parse()

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "ok")
	})

	mux.Handle("/", siteHandler(http.FS(staticFS)))

	log.Printf("curve.thinkwright.ai listening on :%s", *port)

	handler := securityHeaders(mux)
	if err := http.ListenAndServe(":"+*port, handler); err != nil {
		log.Fatal(err)
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// securityHeaders adds standard security headers to all responses.
func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		next.ServeHTTP(w, r)
	})
}

// siteHandler serves the prerendered site from the embedded filesystem.
// Pages are prerendered as <path>.html (SvelteKit trailingSlash "never"), so
// /about serves about.html. Unknown paths get the SvelteKit fallback page
// with a real 404 status, and trailing slashes redirect to the canonical URL.
func siteHandler(fsys http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path

		if p == "/" {
			serveFile(w, r, fsys, "/index.html", http.StatusOK)
			return
		}
		if strings.HasSuffix(p, "/") {
			target := strings.TrimRight(p, "/")
			if r.URL.RawQuery != "" {
				target += "?" + r.URL.RawQuery
			}
			http.Redirect(w, r, target, http.StatusMovedPermanently)
			return
		}

		if isFile(fsys, p) {
			serveFile(w, r, fsys, p, http.StatusOK)
			return
		}
		if path.Ext(p) == "" && isFile(fsys, p+".html") {
			serveFile(w, r, fsys, p+".html", http.StatusOK)
			return
		}

		if isFile(fsys, "/404.html") {
			serveFile(w, r, fsys, "/404.html", http.StatusNotFound)
			return
		}
		http.NotFound(w, r)
	})
}

func isFile(fsys http.FileSystem, name string) bool {
	f, err := fsys.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()
	stat, err := f.Stat()
	return err == nil && !stat.IsDir()
}

// serveFile opens and serves a single file from the embedded filesystem with
// the given status. Non-200 responses are written directly, since
// http.ServeContent always answers 200 (or 206/304).
func serveFile(w http.ResponseWriter, r *http.Request, fsys http.FileSystem, name string, status int) {
	f, err := fsys.Open(name)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	stat, err := f.Stat()
	if err != nil {
		f.Close()
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if stat.IsDir() {
		f.Close()
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	compressible := isCompressible(name)
	if compressible {
		w.Header().Set("Vary", "Accept-Encoding")
	}
	if acceptsGzip(r) && compressible {
		if gz, gzErr := fsys.Open(name + ".gz"); gzErr == nil {
			if gzStat, statErr := gz.Stat(); statErr == nil && !gzStat.IsDir() {
				f.Close()
				f = gz
				stat = gzStat
				w.Header().Set("Content-Encoding", "gzip")
			} else {
				gz.Close()
			}
		}
	}

	// Set Content-Type and Cache-Control based on file path and extension.
	switch {
	case strings.HasSuffix(name, ".html"):
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
	case strings.HasSuffix(name, ".css"):
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
	case strings.HasSuffix(name, ".js"):
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	case strings.HasSuffix(name, ".svg"):
		w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
	case strings.HasSuffix(name, ".png"):
		w.Header().Set("Content-Type", "image/png")
	case strings.HasSuffix(name, ".xml"):
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
	case strings.HasSuffix(name, ".txt"):
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	case strings.HasSuffix(name, ".wasm"):
		w.Header().Set("Content-Type", "application/wasm")
	case strings.HasSuffix(name, ".json"):
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	case strings.HasSuffix(name, ".parquet"):
		w.Header().Set("Content-Type", "application/vnd.apache.parquet")
	}
	w.Header().Set("Cache-Control", cacheControl(name))

	defer f.Close()
	if status != http.StatusOK {
		w.Header().Set("Cache-Control", "no-cache")
		w.WriteHeader(status)
		if r.Method != http.MethodHead {
			io.Copy(w, f)
		}
		return
	}
	seeker, ok := f.(io.ReadSeeker)
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, name, stat.ModTime(), seeker)
}

func acceptsGzip(r *http.Request) bool {
	for _, part := range strings.Split(r.Header.Get("Accept-Encoding"), ",") {
		token := strings.TrimSpace(strings.Split(part, ";")[0])
		if strings.EqualFold(token, "gzip") && !strings.Contains(strings.ReplaceAll(part, " ", ""), ";q=0") {
			return true
		}
	}
	return false
}

func isCompressible(name string) bool {
	switch {
	case strings.HasSuffix(name, ".html"),
		strings.HasSuffix(name, ".js"),
		strings.HasSuffix(name, ".css"),
		strings.HasSuffix(name, ".json"),
		strings.HasSuffix(name, ".svg"),
		strings.HasSuffix(name, ".xml"),
		strings.HasSuffix(name, ".txt"),
		strings.HasSuffix(name, ".wasm"):
		return true
	default:
		return false
	}
}

func cacheControl(name string) string {
	switch {
	case strings.HasSuffix(name, ".html"):
		return "no-cache"
	case strings.HasPrefix(name, "/_app/immutable/"):
		return "public, max-age=31536000, immutable"
	case strings.HasPrefix(name, "/duckdb/"):
		return "public, max-age=86400"
	case strings.HasPrefix(name, "/data/"):
		return "public, max-age=300, stale-while-revalidate=86400"
	case strings.HasSuffix(name, ".svg"), strings.HasSuffix(name, ".png"):
		return "public, max-age=3600"
	case strings.HasSuffix(name, ".xml"), strings.HasSuffix(name, ".txt"), strings.HasSuffix(name, ".json"):
		return "public, max-age=300"
	default:
		return "no-cache"
	}
}
