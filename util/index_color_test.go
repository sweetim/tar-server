package util

import "testing"

func TestIndexColor(t *testing.T) {
	cases := []struct {
		in       int
		expected string
	}{
		{0, "gray"},
		{1, "blue"},
		{2, "red"},
		{3, "red"},
	}

	for _, c := range cases {
		actual := IndexColor(c.in, "gray", "blue", "red")

		if actual != c.expected {
			t.Errorf("actual (%v) != expected (%v)",
				actual,
				c.expected)
		}
	}
}
