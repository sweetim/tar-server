package util

import (
	"testing"
)

func TestUnitSuffix(t *testing.T) {
	cases := []struct {
		in    int64
		text  string
		power int
	}{
		{0, "0 B", 0},
		{321, "321 B", 0},
		{654321, "639 kB", 1},
		{987654321, "942 MB", 2},
		{123456789012, "115 GB", 3},
		{123456789012345, "114978 GB", 3},
	}

	for _, c := range cases {
		actual := UnitSuffix(c.in)

		if actual.Text != c.text &&
			actual.Power != c.power {
			t.Errorf("text : actual (%v) != expected (%v)\npower: actual (%v) != expected (%v)",
				actual.Text,
				c.text,
				actual.Power,
				c.power)
		}
	}
}
