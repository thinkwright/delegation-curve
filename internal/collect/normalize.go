package collect

import "math"

// NormMethod determines how a raw value is mapped to 0-100.
type NormMethod int

const (
	DirectPercent NormMethod = iota // value is already 0-100
	LinearClamp                     // linear interpolation between Min and Max
	LogScale                        // log-scale for values spanning orders of magnitude
)

// NormConfig defines how to normalize a raw value to the 0-100 scale.
type NormConfig struct {
	Method NormMethod
	Min    float64 // LinearClamp: value at Min → 0
	Max    float64 // LinearClamp: value at Max → 100
	LogMin float64 // LogScale: log10(min expected) → 0
	LogMax float64 // LogScale: log10(max expected) → 100
}

// Normalize converts a raw value into a 0-100 normalized score.
// Returns 0 for NaN or Inf inputs.
func Normalize(raw float64, cfg NormConfig) float64 {
	if math.IsNaN(raw) || math.IsInf(raw, 0) {
		return 0
	}
	switch cfg.Method {
	case DirectPercent:
		return clamp(raw, 0, 100)
	case LinearClamp:
		if cfg.Max == cfg.Min {
			return 0
		}
		return clamp((raw-cfg.Min)/(cfg.Max-cfg.Min)*100, 0, 100)
	case LogScale:
		if raw <= 0 {
			return 0
		}
		logVal := math.Log10(raw)
		if cfg.LogMax == cfg.LogMin {
			return 0
		}
		return clamp((logVal-cfg.LogMin)/(cfg.LogMax-cfg.LogMin)*100, 0, 100)
	}
	return 0
}

func clamp(v, lo, hi float64) float64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
