package main

import (
	"fmt"
	"os"
)

func ValidArg(str string) bool {
	for i := range str {
		if !((str[i] >= '1' && str[i] <= '9') || str[i] == '.') {
			return false
		}
	}
	return true
}

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

func SolveSud(result *[9][9]int) bool {
	if !HasEmptyCell(*result) {
		return true
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if result[i][j] == 0 {
				for k := 1; k <= 9; k++ {
					result[i][j] = k
					if IsValidSud(*result) {
						if SolveSud(result) {
							return true
						}
						result[i][j] = 0
					} else {
						result[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return false
}

func HasEmptyCell(result [9][9]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if result[i][j] == 0 {
				return true
			}
		}
	}
	return false
}

func HasDuplicates(counter [10]int) bool {
	for i := 1; i < 10; i++ {
		if counter[i] > 1 {
			return true
		}
	}
	return false
}

func IsValidSud(result [9][9]int) bool {
	//check duplicates by row
	for i := 0; i < 9; i++ {
		counter := [10]int{}
		counter2 := [10]int{}
		for j := 0; j < 9; j++ {
			counter[result[i][j]]++
			counter2[result[j][i]]++
		}
		if HasDuplicates(counter) {
			return false
		}
		if HasDuplicates(counter2) {
			return false
		}
	}

	//check 3x3 section
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			counter := [10]int{}
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					counter[result[row][col]]++
				}
				if HasDuplicates(counter) {
					return false
				}
			}
		}
	}

	return true
}

func main() {
	args := os.Args
	lenArgs := len(args)
	if lenArgs != 10 {
		fmt.Println("Error")
	} else {
		var result [9][9]int
		if ParseArgs(args, &result) {
			if SolveSud(&result) {
				for i := 0; i < 9; i++ {
					for j := 0; j < 9; j++ {
						fmt.Print(result[i][j])
						if j != 8 {
							fmt.Print(" ")
						}
					}
					fmt.Println()
				}
			}
		} else {
			fmt.Println("Error")
		}
	}
}
