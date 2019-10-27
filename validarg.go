package piscine

func ValidArg(str string) bool {
	for i := range str {
		if !((str[i] >= '1' && str[i] <= '9') || str[i] == '.') {
			return false
		}
	}
	return true
}
