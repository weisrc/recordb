package tree

func compress(char byte) byte {
	if char >= 97 && char <= 122 {
		return char - 97
	} else if char == 46 {
		return 36
	} else if char == 45 {
		return 37
	} else if char >= 48 && char <= 57 {
		return char - 48 + 26
	} else if char >= 65 && char <= 90 {
		return char - 65
	}
	return 37
}
