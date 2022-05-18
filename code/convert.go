package code

func FromID(id uint) uint {
	if id <= 65 {
		return id
	} else if id >= 99 && id <= 109 {
		return id - 99 + 66
	} else if id >= 249 && id <= 259 {
		return id - 249 + 77
	} else if id == 32_768 {
		return TA
	} else if id == 32_769 {
		return DLV
	}
	return ALL
}

func ToID(code uint) uint {
	if code == TA {
		return 32_768
	} else if code == DLV {
		return 32_769
	} else if code <= 65 {
		return code
	} else if code <= 76 {
		return code - 66 + 99
	} else if code <= 86 {
		return code - 77 + 249
	}
	return 0
}
