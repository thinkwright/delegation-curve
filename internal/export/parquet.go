package export

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/parquet-go/parquet-go"
	"github.com/parquet-go/parquet-go/compress/zstd"
)

// WriteTable writes a slice of typed rows to a Parquet file with zstd compression.
func WriteTable[T any](dir, name string, rows []T) error {
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("create dir %s: %w", dir, err)
	}
	path := filepath.Join(dir, name+".parquet")

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create file %s: %w", path, err)
	}
	defer f.Close()

	writer := parquet.NewGenericWriter[T](f, parquet.Compression(&zstd.Codec{}))

	if _, err := writer.Write(rows); err != nil {
		return fmt.Errorf("write rows to %s: %w", name, err)
	}
	if err := writer.Close(); err != nil {
		return fmt.Errorf("close writer for %s: %w", name, err)
	}
	if err := f.Sync(); err != nil {
		return fmt.Errorf("sync file %s: %w", name, err)
	}

	stat, _ := os.Stat(path)
	fmt.Printf("  %-25s %d rows  %s\n", name+".parquet", len(rows), humanSize(stat.Size()))
	return nil
}

func humanSize(b int64) string {
	switch {
	case b >= 1<<20:
		return fmt.Sprintf("%.1f MB", float64(b)/float64(1<<20))
	case b >= 1<<10:
		return fmt.Sprintf("%.1f KB", float64(b)/float64(1<<10))
	default:
		return fmt.Sprintf("%d B", b)
	}
}
