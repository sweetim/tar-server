package util

import (
	"fmt"
	"math"
)

// UnitSuffixResult provides the formatted text and base power
type UnitSuffixResult struct {
	Text  string
	Power int
}

var units = [...]string{"B", "kB", "MB", "GB"}

// UnitSuffix will suffix the value with unit
func UnitSuffix(input int64) UnitSuffixResult {
	v := float64(input)

	if v == 0 {
		return UnitSuffixResult{
			"0 B",
			0,
		}
	}

	power := int(math.Floor(math.Log10(v) / math.Log10(1024)))
	if power >= len(units) {
		power = len(units) - 1
	}

	base := v / math.Pow(1024, float64(power))

	return UnitSuffixResult{
		fmt.Sprintf("%.f %v", base, units[power]),
		power,
	}
}
