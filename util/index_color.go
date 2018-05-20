package util

// IndexColor match index with color input
func IndexColor(index int, colors ...string) string {
	if index < len(colors) {
		return colors[index]
	}

	return colors[len(colors)-1]
}
