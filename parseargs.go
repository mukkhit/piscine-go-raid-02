package piscine

func ParseArgs(args []string, result *[9][9]int) bool {
	for i := 1; i < len(args); i++ {
		if len(args[i]) != 9 {
			return false
		} else {
			if ValidArg(args[i]) {
				for j := 0; j < 9; j++ {
					if args[i][j] != '.' {
						result[i-1][j] = int(args[i][j] - 48)
					}
				}
			} else {
				return false
			}
		}
	}
	return true
}
