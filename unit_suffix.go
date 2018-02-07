package main

import (
	"fmt"
	"math"
)

// UnitSuffix will suffix the value with unit
func UnitSuffix(v float64) string {
	var units = [...]string{"B", "kB", "MB", "GB"}

	if v == 0 {
		return "0 B"
	}

	power := int(math.Floor(math.Log10(v) / math.Log10(1024)))
	if power >= len(units) {
		power = len(units) - 1
	}

	base := v / math.Pow(1024, float64(power))

	return fmt.Sprintf("%.f %v", base, units[power])
}
