package main

//import "fmt"
import "github.com/01-edu/z01"

func main() {
	args := []string{"main", ".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	var massive [9][9][11]rune
	//vvod parametrov
	for y := 1; y <= 9; y++ {
		slovo := args[y] //".96.4...1"
		for x := range slovo {
			massive[y-1][x][0] = rune(slovo[x])
			if rune(slovo[x]) == '.' {
				num := 1
				for z := '1'; z <= '9'; z++ {
					massive[y-1][x][num] = z
					num++
				}
				massive[y-1][x][10] = '9'
			}
		}
	}
	massive2 := massive
	z01.PrintRune(10)
	for y := 0; y <= 8; y++ {
		for x := 0; x <= 8; x++ {
			z01.PrintRune(massive2[y][x][0])
		}
		z01.PrintRune(':')
		for x := 0; x <= 8; x++ {
			for z := 1; z <= 10; z++ {
				z01.PrintRune(massive2[y][x][z])
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
	//poisk resheniya
	//massive2 = possible1(massive)
	//neizvestno:=1
	for i := 1; neizvestno > 0; i++ {
		massive2 = possible2(massive2)
		z01.PrintRune(10)
		for y := 0; y <= 8; y++ {
			for x := 0; x <= 8; x++ {
				z01.PrintRune(massive2[y][x][0])
			}
			z01.PrintRune(':')
			for x := 0; x <= 8; x++ {
				for z := 1; z <= 10; z++ {
					z01.PrintRune(massive2[y][x][z])
				}
				z01.PrintRune(' ')
			}
			z01.PrintRune('\n')
		}
		massive2 = possible1(massive2)
		z01.PrintRune(10)
		for y := 0; y <= 8; y++ {
			for x := 0; x <= 8; x++ {
				z01.PrintRune(massive2[y][x][0])
			}
			z01.PrintRune(':')
			for x := 0; x <= 8; x++ {
				for z := 1; z <= 10; z++ {
					z01.PrintRune(massive2[y][x][z])
				}
				z01.PrintRune(' ')
			}
			z01.PrintRune('\n')
		}
		neizvestno := 0
		for y := 0; y <= 8; y++ {
			for x := 0; x <= 8; x++ {
				if massive[y][x][0] == '.' {
					neizvestno++
				}
			}
		}
	}
	//massive2 = possible(massive2)
	//massive2 = possible(massive2)
	// print massive
	//massive2:=massive
	z01.PrintRune(10)
	for y := 0; y <= 8; y++ {
		for x := 0; x <= 8; x++ {
			z01.PrintRune(massive2[y][x][0])
		}
		z01.PrintRune(':')
		for x := 0; x <= 8; x++ {
			for z := 1; z <= 10; z++ {
				z01.PrintRune(massive2[y][x][z])
			}
			z01.PrintRune(' ')
		}
		z01.PrintRune('\n')
	}
}

func possible1(massive [9][9][11]rune) [9][9][11]rune {
	var massive2 [9][9][11]rune
	massive2 = massive
	nashli := false
	//z01.PrintRune('\n')
	//if massive2[0][0][8] == '!' {
	//	z01.PrintRune(massive2[0][0][8])
	//	massive2[0][0][8] = '.'
	//	z01.PrintRune(massive2[0][0][8])
	//}
	for y := 0; y <= 8; y++ {
		for x := 0; x <= 8; x++ {
			for ryad := 0; ryad <= 8; ryad++ { //idu po ryadu - isklychayu cifry ryada
				if massive2[y][ryad][0] != '.' { //0  0
					for z := 1; z <= 9; z++ {
						if massive2[y][x][z] == massive2[y][ryad][0] {
							massive2[y][x][z] = '.'
							massive2[y][x][10]-- //8
						}
					}
				} //if not .
			} //for ryad
			for col := 0; col <= 8; col++ { //idu po stolbcu - isklychayu cifry stolbca
				if massive2[col][x][0] != '.' { //0  0
					for z := 1; z <= 9; z++ {
						if massive2[y][x][z] == massive2[col][x][0] {
							massive2[y][x][z] = '.'
							massive2[y][x][10]-- //8
						}
					}
				} //if not .
			} //for col
			for square1 := (y / 3) * 3; square1 <= (y/3)*3+2; square1++ { //idu po square - isklychayu cifry square
				for square2 := (x / 3) * 3; square2 <= (x/3)*3+2; square2++ {
					if massive2[square1][square2][0] != '.' { //0  0
						for z := 1; z <= 9; z++ {
							if massive2[y][x][z] == massive2[square1][square2][0] {
								massive2[y][x][z] = '.'
								massive2[y][x][10]-- //8
							}
						}
					} //if not .
				} //for square2
			} //for square1
			//my ubrali vse nevozmojnye znachenia
			if massive2[y][x][10] == '1' {
				for z := 1; z <= 9; z++ {
					if massive2[y][x][z] != '.' {
						massive2[y][x][0] = massive2[y][x][z]
						massive2[y][x][z] = '.'
						massive2[y][x][10] = '.'
						nashli = true
						break
					}
				}
			}
			if nashli == true {
				break
			}
		}
		if nashli == true {
			//massive2[0][0][8] = '!'
			return massive2
		}
	}
	return massive2
}

func possible2(massive [9][9][11]rune) [9][9][11]rune {

	//-----------------cifry unikalnye
	massive2 := massive
	nashli := false
	nashli = nashli
	var numbers [2][9]rune
	var stolbec [9]int
	for y := 0; y <= 8; y++ {
		numbers = [2][9]rune{
			{'1', '2', '3', '4', '5', '6', '7', '8', '9'},
			{'0', '0', '0', '0', '0', '0', '0', '0', '0'},
		}
		stolbec = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
		for x := 0; x <= 8; x++ {
			if massive2[y][x][0] == '.' {
				for z := 1; z <= 9; z++ {
					if massive2[y][x][z] != '.' {
						for n := 0; n <= 8; n++ {
							if massive2[y][x][z] == numbers[0][n] {
								numbers[1][n]++ //kolichestvo cifr
								stolbec[n] = x  //stolbec yacheiki
							}
						}
					} //if not .
				}
			} //if z not nil
		}
		/*for i := 0; i < 9; i++ {
			z01.PrintRune(numbers[0][i])
			z01.PrintRune('=')
			z01.PrintRune(numbers[1][i])
			z01.PrintRune(';')
		}*/
		for n := 0; n <= 8; n++ {
			if numbers[1][n] == '1' {
				massive2[y][stolbec[n]][0] = numbers[0][n]
				nashli = true
				//massive2[0][0][8] = '!'
				for z := 1; z <= 10; z++ {
					massive2[y][stolbec[n]][z] = '.'
				}
			}
		}
	}
	return massive2
}
