package transform

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

var (
	yearOnly    = regexp.MustCompile(`^(\d{4})$`)
	yearQuarter = regexp.MustCompile(`^(\d{4}) Q([1-4])$|^Q([1-4]) (\d{4})$`)
	yearHalf    = regexp.MustCompile(`^(\d{4}) H([12])$|^H([12]) (\d{4})$`)
)

// PeriodBounds parses a run measurement period label ("2025", "2026 Q2",
// "Q2 2026", "2026 H2") into its first and last day. Unrecognized labels
// fall back to the full measurement year.
func PeriodBounds(period string, year int) (start, end time.Time) {
	months := func(y, firstMonth, n int) (time.Time, time.Time) {
		s := time.Date(y, time.Month(firstMonth), 1, 0, 0, 0, 0, time.UTC)
		return s, s.AddDate(0, n, -1)
	}
	if m := yearOnly.FindStringSubmatch(period); m != nil {
		y, _ := strconv.Atoi(m[1])
		return months(y, 1, 12)
	}
	if m := yearQuarter.FindStringSubmatch(period); m != nil {
		y, q := atoi(m[1], m[4]), atoi(m[2], m[3])
		return months(y, (q-1)*3+1, 3)
	}
	if m := yearHalf.FindStringSubmatch(period); m != nil {
		y, h := atoi(m[1], m[4]), atoi(m[2], m[3])
		return months(y, (h-1)*6+1, 6)
	}
	return months(year, 1, 12)
}

// PeriodMidpoint returns the midpoint of a measurement period as YYYY-MM-DD.
func PeriodMidpoint(period string, year int) string {
	start, end := PeriodBounds(period, year)
	mid := start.Add(end.Sub(start) / 2)
	return mid.Format("2006-01-02")
}

// atoi returns the first non-empty numeric capture.
func atoi(a, b string) int {
	s := a
	if s == "" {
		s = b
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("period regexp captured non-number %q", s))
	}
	return n
}
