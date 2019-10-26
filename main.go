package main

//import "fmt"
import "github.com/01-edu/z01"

func main() {
	args := []string{"main", ".96.4...1", "1...6...4", "5.481.39.", "..795..43", ".3..8....", "4.5.23.18", ".1.63..59", ".59.7.83.", "..359...7"}
	var massive [9][9][11]rune
	//vvod parametrov
	neizvestno := 0
	for y := 1; y <= 9; y++ {
		slovo := args[y] //".96.4...1"
		for x := range slovo {
			massive[y-1][x][0] = rune(slovo[x])
			if rune(slovo[x]) == '.' {
				neizvestno++
			}
			num := 1
			for z := '1'; z <= '9'; z++ {
				massive[y-1][x][num] = z
				num++
			}
			massive[y-1][x][10] = '9'
		}
	}
	//poisk resheniya
	massive2 := possible(massive)
	for i := 1; i <= 10000000000; i++ {
		if massive2[8][8][8] == '!' {
			massive2 = possible(massive2)
		} else {
			massive2 = massive2
			break
		}
	}
	// print massive
	for y := 0; y <= 8; y++ {
		for x := 0; x <= 8; x++ {
			z01.PrintRune(massive2[y][x][0])
		}
		z01.PrintRune('\n')
	}
}

func possible(massive [9][9][11]rune) [9][9][11]rune {
	var massive2 [9][9][11]rune
	massive2 = massive
	nashli := false
	if massive2[8][8][8] == '!' {
		massive2[8][8][8] = '.'
	}
	for y := 0; y <= 8; y++ {
		for x := 0; x <= 8; x++ {
			//v pervoi yacheike
			for ryad := 0; ryad <= 8; ryad++ { //idu po ryadu - isklychayu cifry ryada
				if massive2[y][ryad][0] != '.' { //0  0
					for z := 0; z <= 8; z++ {
						if massive2[y][x][z] == massive2[y][ryad][0] {
							massive2[y][x][z] = '.'
							massive2[y][x][10]-- //8
						}
					}
				} //if not .
			} //for ryad
			for col := 0; col <= 8; col++ { //idu po stolbcu - isklychayu cifry stolbca
				if massive2[col][x][0] != '.' { //0  0
					for z := 0; z <= 8; z++ {
						if massive2[y][x][z] == massive2[col][x][0] {
							massive2[y][x][z] = '.'
							massive2[y][x][10]-- //8
						}
					}
				} //if not .
			} //for col
			/*if y<=2 &&x<=2{
				//2/3=0     4 0 - naoborot         |8  8
				sq1:=0  //4/3  4/3*3                 |8/3 = 2*3
				sq2:=0   //   0/3*3                |8/3 = 2*3
			}*/
			for square1 := (y / 3) * 3; square1 <= (y/3)*3+2; square1++ { //idu po square - isklychayu cifry square
				for square2 := (x / 3) * 3; square2 <= (x/3)*3+2; square2++ {
					if massive2[square1][square2][0] != '.' { //0  0
						for z := 0; z <= 8; z++ {
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
				for z := 0; z <= 8; z++ {
					if massive2[y][x][z] != '.' {
						massive2[y][x][0] = massive2[y][x][z]
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
			massive2[8][8][8] = '!'
			break
		}
	}
	return massive2
}
