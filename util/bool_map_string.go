package util

// BoolMapString map the input boolean to string
func BoolMapString(trueText string, falseText string, input bool) string {
	if input {
		return trueText
	}

	return falseText
}
