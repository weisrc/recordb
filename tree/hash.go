package tree

// range reduction byte [0, 255] -> domain name character [0, 37], custom hash function
func hashChar(char byte) byte {
	if char >= 97 && char <= 122 {
		return char - 97 // lower case alphabet
	} else if char == 46 {
		return 36 // dot
	} else if char == 45 {
		return 37 // hyphen
	} else if char >= 48 && char <= 57 {
		return char - 48 + 26
	} else if char >= 65 && char <= 90 {
		return char - 65 // upper case alphabet
	}
	return 37
}

func unhashChar(char byte) byte {
	if char <= 25 {
		return char + 97 // alphabet
	} else if char >= 26 && char <= 35 {
		return char - 26 + 48 // digit
	} else if char == 36 {
		return 46
	} else {
		return 45
	}
}

func Hash(key string) string {
	length := len(key)
	segment := make([]byte, length)
	for i := 0; i < length; i++ {
		segment[i] = hashChar(key[i])
	}
	return string(segment)
}

func Unhash(segment string) string {
	length := len(segment)
	key := make([]byte, length)
	for i := 0; i < length; i++ {
		key[i] = unhashChar(segment[i])
	}
	return string(key)
}
