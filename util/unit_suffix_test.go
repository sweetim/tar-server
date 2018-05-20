package util

import (
	"testing"
)

func TestUnitSuffix(t *testing.T) {
	cases := []struct {
		in       int64
		expected string
	}{
		{0, "0 B"},
		{321, "321 B"},
		{654321, "639 kB"},
		{987654321, "942 MB"},
		{123456789012, "115 GB"},
		{123456789012345, "114978 GB"},
	}

	for _, c := range cases {
		actual := UnitSuffix(c.in)

		if actual != c.expected {
			t.Errorf("actual (%v) != expected (%v)",
				actual,
				c.expected)
		}
	}
}
