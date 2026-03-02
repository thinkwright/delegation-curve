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

	mux.Handle("/", spaHandler(http.FS(staticFS)))

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

// spaHandler serves static files from the embedded filesystem.
// If a file exists, it is served directly. Otherwise, index.html is served
// as a fallback for SvelteKit client-side routing.
func spaHandler(fsys http.FileSystem) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p := r.URL.Path

		// Root serves landing page.
		if p == "/" {
			serveFile(w, r, fsys, "/index.html")
			return
		}

		// Try the requested path as a file.
		f, err := fsys.Open(p)
		if err == nil {
			defer f.Close()
			stat, _ := f.Stat()
			if !stat.IsDir() {
				serveFile(w, r, fsys, p)
				return
			}
			// Directory — try index.html inside it.
			idx := path.Join(p, "index.html")
			if fi, err := fsys.Open(idx); err == nil {
				fi.Close()
				serveFile(w, r, fsys, idx)
				return
			}
		}

		// SPA fallback: serve index.html for client-side routing.
		serveFile(w, r, fsys, "/index.html")
	})
}

// serveFile opens and serves a single file from the embedded filesystem.
func serveFile(w http.ResponseWriter, r *http.Request, fsys http.FileSystem, name string) {
	f, err := fsys.Open(name)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if stat.IsDir() {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	// Set Content-Type and Cache-Control based on file extension.
	switch {
	case strings.HasSuffix(name, ".html"):
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
	case strings.HasSuffix(name, ".css"):
		w.Header().Set("Content-Type", "text/css; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
	case strings.HasSuffix(name, ".js"):
		w.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		w.Header().Set("Cache-Control", "no-cache")
	case strings.HasSuffix(name, ".svg"):
		w.Header().Set("Content-Type", "image/svg+xml; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=300")
	case strings.HasSuffix(name, ".png"):
		w.Header().Set("Content-Type", "image/png")
		w.Header().Set("Cache-Control", "public, max-age=3600")
	case strings.HasSuffix(name, ".xml"):
		w.Header().Set("Content-Type", "application/xml; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=300")
	case strings.HasSuffix(name, ".txt"):
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=300")
	case strings.HasSuffix(name, ".wasm"):
		w.Header().Set("Content-Type", "application/wasm")
		w.Header().Set("Cache-Control", "public, max-age=3600")
	case strings.HasSuffix(name, ".json"):
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Cache-Control", "public, max-age=300")
	case strings.HasSuffix(name, ".parquet"):
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Cache-Control", "no-cache")
	}

	seeker, ok := f.(io.ReadSeeker)
	if !ok {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	http.ServeContent(w, r, name, stat.ModTime(), seeker)
}
