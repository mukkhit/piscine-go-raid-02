package main 

// import "fmt"
import "github.com/01-edu/z01"

func main2 () {
	args:=[]string{"main", ".96.4...1" ,"1...6...4" ,"5.481.39." ,"..795..43", ".3..8...." ,"4.5.23.18", ".1.63..59" ,".59.7.83." ,"..359...7"}
	var massive [9][9]rune
	//vvod parametrov
	for y := 1; y <= 9; y++ {
		slovo:= args[y] //".96.4...1"
		for x:=range slovo{
			massive[y-1][x]= rune(slovo[x])
			//z01.PrintRune(massive[y-1][x])
		}
		//z01.PrintRune('\n')
	}
	//poisk resheniya
	var tochki[] int
	itochki:=0
	for y := 0; y <= 8; y++ {
		for x:=0; x<=8;x++{
			if massive[y][x]=='.'{
				tochki[itochki]=1
				for z := '1'; z <= '9'; z++ {
					massive[y][x] = z	
					break
				}
			}
		}
	}
	// print massive
	for y := 0; y <= 8; y++ {
		for x:=0; x<=8;x++{
			z01.PrintRune(massive[y][x])
		}
		z01.PrintRune('\n')
	}
}