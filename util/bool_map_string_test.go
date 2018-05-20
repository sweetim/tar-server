package util

import "testing"

func TestBoolMapString(t *testing.T) {
	cases := []struct {
		text     []string
		in       bool
		expected string
	}{
		{
			[]string{
				"folder",
				"insert_drive_file",
			},
			true,
			"folder",
		},
		{
			[]string{
				"folder",
				"insert_drive_file",
			},
			false,
			"insert_drive_file",
		},
	}

	for _, c := range cases {
		actual := BoolMapString(
			c.text[0],
			c.text[1],
			c.in)

		if actual != c.expected {
			t.Errorf("actual (%v) != expected (%v)",
				actual,
				c.expected)
		}
	}
}
