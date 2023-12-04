package main

import (
	"unicode"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	// start := time.Now()
	main1()
	// end := time.Now()
	// println(end.Sub(start).Milliseconds())
}

type symbol struct {
	symbol rune
	v      tj.Vec2
}

func symbolcheck(grid []string, x_start int, x_end int, row int, gears map[tj.Vec2]int) symbol {
	// println()
	y_start := max(0, row-1)
	y_end := min(len(grid)-1, row+1)
	x_start = max(0, x_start-1)
	x_end = min(len(grid[0])-1, x_end)

	for y := y_start; y <= y_end; y++ {
		for x := x_start; x <= x_end; x++ {
			// print(string(grid[y][x]))
			if grid[y][x] != '.' && !unicode.IsNumber(rune(grid[y][x])) {

				return symbol{symbol: rune(grid[y][x]), v: tj.Vec2{X: x, Y: y}}
			}
		}
	}
	return symbol{symbol: '.', v: tj.Vec2{}}
}

func main1() {
	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DIntSlice("input", ' ')

	total := 0
	gears := make(map[tj.Vec2]int)
	ratiocount := make(map[tj.Vec2]int)
	// count := 0
	for j, line := range b {
		numberstart := -1
		numberend := -1
		for i, char := range line {
			// print(string(char))
			// print(numberstart)
			// print(" ", unicode.IsNumber(char), " ")
			if unicode.IsNumber(char) {
				if numberstart == -1 {
					numberstart = i
				} else if i == len(line)-1 {
					numberend = i + 1
				}
			} else {
				if numberstart != -1 {
					if !unicode.IsNumber(char) {
						numberend = i
					}
				}
			}
			if numberstart != -1 && numberend != -1 {
				sym := symbolcheck(b, numberstart, i, j, gears)
				if sym.symbol != '.' {
					num := tj.Str2int(line[numberstart:numberend])
					total += num
					if sym.symbol == '*' {
						// println("Found gear, ratio", num)
						ratiocount[sym.v] += 1
						if gears[sym.v] == 0 {
							gears[sym.v] = num
						} else {
							gears[sym.v] *= num

						}
					}
				}
				numberstart = -1
				numberend = -1
			}
		}
	}
	// 1192 numbers
	// println()
	// println(count)
	ratiototal := 0
	for k, v := range gears {
		if ratiocount[k] == 2 {
			ratiototal += v
		}
	}
	println(total)
	println(ratiototal)
	// 524316 wrong
	// 518219 wrong
	// 518665 wrong
	// 524762 wrong
	// 522726 right

}
