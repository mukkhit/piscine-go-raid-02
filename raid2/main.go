package main

import (
	piscine ".."
	"fmt"
	"os"
)

func main() {
	args := os.Args
	lenArgs := len(args)
	if lenArgs != 10 {
		fmt.Println("Error")
	} else {
		var result [9][9]int
		if piscine.ParseArgs(args, &result) {
			for i := 0; i < 9; i++ {
				for j := 0; j < 9; j++ {
					fmt.Print(result[i][j])
					fmt.Print(" ")
				}
				fmt.Println()
			}
		} else {
			fmt.Println("Error")
		}
	}
}

// func parseInput(input string) [9][9]int {
// 	board := [9][9]int{}
// 	//vvod parametrov
// 	for y := 1; y <= 9; y++ {
// 		slovo := args[y] //".96.4...1"
// 		for x := range slovo {
// 			board[y-1][x][0] = rune(slovo[x])
// 			if rune(slovo[x]) == '.' {
// 				num := 1
// 				for z := '1'; z <= '9'; z++ {
// 					board[y-1][x][num] = z
// 					num++
// 				}
// 				board[y-1][x][10] = '9'
// 			}
// 		}
// 	}
// 	/* massive2 := massive
// 	z01.PrintRune(10)
// 	for y := 0; y <= 8; y++ {
// 		for x := 0; x <= 8; x++ {
// 			z01.PrintRune(massive2[y][x][0])
// 		}
// 		z01.PrintRune(':')
// 		for x := 0; x <= 8; x++ {
// 			for z := 1; z <= 10; z++ {
// 				z01.PrintRune(massive2[y][x][z])
// 			}
// 			z01.PrintRune(' ')
// 		}
// 		z01.PrintRune('\n')
// 	} */
// 	return board
// }

// func isBoardValid(board *[9][9]int) bool {

// 	//check duplicates by row
// 	for row := 0; row < 9; row++ {
// 		counter := [10]int{}
// 		for col := 0; col < 9; col++ {
// 			counter[board[row][col]]++
// 		}
// 		if hasDuplicates(counter) {
// 			return false
// 		}
// 	}

// 	//check duplicates by column
// 	for row := 0; row < 9; row++ {
// 		counter := [10]int{}
// 		for col := 0; col < 9; col++ {
// 			counter[board[col][row]]++
// 		}
// 		if hasDuplicates(counter) {
// 			return false
// 		}
// 	}

// 	//check 3x3 section
// 	for i := 0; i < 9; i += 3 {
// 		for j := 0; j < 9; j += 3 {
// 			counter := [10]int{}
// 			for row := i; row < i+3; row++ {
// 				for col := j; col < j+3; col++ {
// 					counter[board[row][col]]++
// 				}
// 				if hasDuplicates(counter) {
// 					return false
// 				}
// 			}
// 		}
// 	}

// 	return true
// }

// func backtrack(board *[9][9]int) bool {
// 	if !hasEmptyCell(board) {
// 		return true
// 	}
// 	for i := 0; i < 9; i++ {
// 		for j := 0; j < 9; j++ {
// 			if board[i][j] == 0 {
// 				for candidate := 9; candidate >= 1; candidate-- {
// 					board[i][j] = candidate
// 					if isBoardValid(board) {
// 						if backtrack(board) {
// 							return true
// 						}
// 						board[i][j] = 0
// 					} else {
// 						board[i][j] = 0
// 					}
// 				}
// 				return false
// 			}
// 		}
// 	}
// 	return false
// }

// func printBoard(board [9][9]int) {
// 	fmt.Println("+-------+-------+-------+")
// 	for row := 0; row < 9; row++ {
// 		fmt.Print("| ")
// 		for col := 0; col < 9; col++ {
// 			if col == 3 || col == 6 {
// 				fmt.Print("| ")
// 			}
// 			fmt.Printf("%d ", board[row][col])
// 			if col == 8 {
// 				fmt.Print("|")
// 			}
// 		}
// 		if row == 2 || row == 5 || row == 8 {
// 			fmt.Println("\n+-------+-------+-------+")
// 		} else {
// 			fmt.Println()
// 		}
// 	}
// }
