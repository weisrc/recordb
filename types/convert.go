package types

func Hash(id uint16) uint16 {
	if id <= 65 {
		return id
	} else if id >= 99 && id <= 109 {
		return id - 99 + 66
	} else if id >= 249 && id <= 259 {
		return id - 249 + 77
	} else if id == 32_768 {
		return TA
	} else if id == 32_769 {
		return 54
	}
	return 83
}
