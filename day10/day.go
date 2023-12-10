package main

import (
	"flag"
	"fmt"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	inputFlag := flag.Bool("f", false, "use the final input")
	flag.Parse()

	input := ""

	if *inputFlag {
		input = "input_real"
	} else {
		input = "input"
	}

	// start := time.Now()
	// main1(input)
	main2(input)
	// main2dumb(input)
	// end := time.Now()
	// println(end.Sub(start).Milliseconds())
}

type pipe struct {
	r rune
}

func PrintMap(pipeMap [][]rune) {
	for _, line := range pipeMap {
		for _, p := range line {
			print(string(p))
		}
		println()
	}
}

func PrintMapNumbers(pipeMap [][]int) {
	for _, line := range pipeMap {
		for _, p := range line {
			fmt.Printf("%5d", p)
		}
		println()
	}
}

func GetDirections(r rune) []tj.Vec2 {
	switch r {
	case '|':
		return []tj.Vec2{tj.Vec2{X: 0, Y: -1}, tj.Vec2{X: 0, Y: 1}}
	case '-':
		return []tj.Vec2{tj.Vec2{X: -1, Y: 0}, tj.Vec2{X: 1, Y: 0}}
	case 'L':
		return []tj.Vec2{tj.Vec2{X: 0, Y: -1}, tj.Vec2{X: 1, Y: 0}}
	case 'J':
		return []tj.Vec2{tj.Vec2{X: 0, Y: -1}, tj.Vec2{X: -1, Y: 0}}
	case '7':
		return []tj.Vec2{tj.Vec2{X: 0, Y: 1}, tj.Vec2{X: -1, Y: 0}}
	case 'F':
		return []tj.Vec2{tj.Vec2{X: 0, Y: 1}, tj.Vec2{X: 1, Y: 0}}
	default:
		return []tj.Vec2{}
	}
}

func DoubleSpace[k rune | int](in [][]k, filler k) [][]k {
	out := [][]k{}
	for y := 0; y < len(in)*2; y++ {
		out = append(out, []k{})
		for x := 0; x < len(in[0])*2; x++ {
			out[y] = append(out[y], filler)
			if y%2 == 0 && x%2 == 0 {
				out[y][x] = in[y/2][x/2]
			}
		}
	}
	return out
}

func CanFloodToEdge(pipeMap [][]int, start tj.Vec2) bool {
	cardinalDirections := [4]tj.Vec2{tj.Vec2{X: 0, Y: -1}, tj.Vec2{X: 0, Y: 1}, tj.Vec2{X: -1, Y: 0}, tj.Vec2{X: 1, Y: 0}}
	// This returns true if the start point can flood to the edge of the map
	// only hitting -1 cells
	if pipeMap[start.Y][start.X] != -1 {
		return false
	}
	pipeMap[start.Y][start.X] = 0
	for {
		added := false
		for y := 0; y < len(pipeMap); y++ {
			for x := 0; x < len(pipeMap[y]); x++ {
				if pipeMap[y][x] == 0 {
					for _, dir := range cardinalDirections {
						if y+dir.Y < 0 || y+dir.Y >= len(pipeMap) || x+dir.X < 0 || x+dir.X >= len(pipeMap[y]) {
							return true
						}
						if pipeMap[y+dir.Y][x+dir.X] == -1 {
							pipeMap[y+dir.Y][x+dir.X] = 0
							added = true
						}
					}
				}
			}
		}
		// PrintMapNumbers(pipeMap)
		// println()
		if !added {
			break
		}
	}
	return false
}

func Copy2dSlice(in [][]int) [][]int {
	out := [][]int{}
	for _, line := range in {
		out = append(out, []int{})
		for _, i := range line {
			out[len(out)-1] = append(out[len(out)-1], i)
		}
	}
	return out
}

func main2(inputfile string) {
	b := tj.FileToSlice(inputfile)

	pipeMap := [][]rune{}

	pipeCountings := [][]int{}

	// Harcode this for now
	startPoint := tj.Vec2{X: 0, Y: 2}

	for _, line := range b {
		pipeMap = append(pipeMap, []rune{})
		// pipeCountings = append(pipeCountings, []int{})
		for _, r := range line {
			if r == 'S' {
				r = '0'
				startPoint = tj.Vec2{X: len(pipeMap[len(pipeMap)-1]), Y: len(pipeMap) - 1}
				// pipeCountings[len(pipeCountings)-1] = append(pipeCountings[len(pipeCountings)-1], 0)
			} else {
				// pipeCountings[len(pipeCountings)-1] = append(pipeCountings[len(pipeCountings)-1], -1)
			}
			pipeMap[len(pipeMap)-1] = append(pipeMap[len(pipeMap)-1], r)
		}
	}
	if inputfile == "input" {
		// Dirty hack. Manually look up the correct rune for the puzzle start point*
		// pipeMap[startPoint.Y][startPoint.X] = 'F'
		pipeMap[startPoint.Y][startPoint.X] = 'F'
	} else {
		pipeMap[startPoint.Y][startPoint.X] = 'L'
	}

	pipeMapDoubled := DoubleSpace(pipeMap, '.')
	// Now fill in the spots between the existing pipes with only - or |
	for y := 0; y < len(pipeMapDoubled); y++ {
		pipeCountings = append(pipeCountings, []int{})
		for x := 0; x < len(pipeMapDoubled[y]); x++ {
			pipeCountings[len(pipeCountings)-1] = append(pipeCountings[len(pipeCountings)-1], -1)
			if x%2 == 0 && y%2 == 0 {
				if startPoint.X == x/2 && startPoint.Y == y/2 {
					pipeCountings[y][x] = 0
				}
			}

			if pipeMapDoubled[y][x] == '.' {
				continue
			}
			dir := GetDirections(pipeMapDoubled[y][x])
			for _, d := range dir {
				pipe := '.'
				if d.X != 0 {
					pipe = '-'
				} else {
					pipe = '|'
				}
				if y+d.Y < 0 || y+d.Y >= len(pipeMapDoubled) || x+d.X < 0 || x+d.X >= len(pipeMapDoubled[y]) {
					continue
				}
				if pipeMapDoubled[y+d.Y][x+d.X] != '.' {
					continue
				}
				pipeMapDoubled[y+d.Y][x+d.X] = pipe
			}
		}
	}
	// PrintMap(pipeMapDoubled)
	// println()
	// PrintMapNumbers(pipeCountings)

	highest := 0
	for {
		added := false
		for y := 0; y < len(pipeMapDoubled); y++ {
			for x := 0; x < len(pipeMapDoubled[y]); x++ {
				if pipeCountings[y][x] >= 0 {
					for _, dir := range GetDirections(pipeMapDoubled[y][x]) {
						if pipeCountings[y+dir.Y][x+dir.X] != -1 {
							continue
						}
						pipeCountings[y+dir.Y][x+dir.X] = pipeCountings[y][x] + 1
						added = true
						if pipeCountings[y+dir.Y][x+dir.X] > highest {
							highest = pipeCountings[y+dir.Y][x+dir.X]
						}
					}
				}
			}
		}
		if !added {
			break
		}
	}

	contained := 0
	for y := 0; y < len(pipeCountings); y++ {
		for x := 0; x < len(pipeCountings[y]); x++ {
			if pipeCountings[y][x] >= 0 {
				pipeCountings[y][x] = 1
			}
		}
	}
	for y := 0; y < len(pipeCountings); y++ {
		for x := 0; x < len(pipeCountings[y]); x++ {
			if pipeCountings[y][x] != -1 {
				continue
			}
			pc := Copy2dSlice(pipeCountings)
			// println()
			if !CanFloodToEdge(pc, tj.Vec2{X: x, Y: y}) {
				// return
				if x%2 == 0 && y%2 == 0 {
					contained += 1
				}
			}
		}
	}

	// PrintMap(pipeMap)
	// pipeMap = DoubleSpace(pipeMap)
	// pipeCountings = DoubleSpace(pipeCountings)
	// PrintMap(pipeMap)

	// PrintMapNumbers(pipeCountings)

	println(highest)
	println(contained)
}

func main1(inputfile string) {
	b := tj.FileToSlice(inputfile)

	pipeMap := [][]rune{}

	pipeCountings := [][]int{}

	// Harcode this for now
	startPoint := tj.Vec2{X: 0, Y: 2}

	for _, line := range b {
		pipeMap = append(pipeMap, []rune{})
		pipeCountings = append(pipeCountings, []int{})
		for _, r := range line {
			if r == 'S' {
				r = '0'
				startPoint = tj.Vec2{X: len(pipeMap[len(pipeMap)-1]), Y: len(pipeMap) - 1}
				pipeCountings[len(pipeCountings)-1] = append(pipeCountings[len(pipeCountings)-1], 0)
			} else {
				pipeCountings[len(pipeCountings)-1] = append(pipeCountings[len(pipeCountings)-1], -1)
			}
			pipeMap[len(pipeMap)-1] = append(pipeMap[len(pipeMap)-1], r)
		}
	}
	if inputfile == "input" {
		// Dirty hack. Manually look up the correct rune for the puzzle start point*
		// pipeMap[startPoint.Y][startPoint.X] = 'F'
		pipeMap[startPoint.Y][startPoint.X] = 'F'
	} else {
		pipeMap[startPoint.Y][startPoint.X] = 'L'
	}

	highest := 0
	for {
		added := false
		for y := 0; y < len(pipeMap); y++ {
			for x := 0; x < len(pipeMap[y]); x++ {
				if pipeCountings[y][x] >= 0 {
					for _, dir := range GetDirections(pipeMap[y][x]) {
						if pipeCountings[y+dir.Y][x+dir.X] != -1 {
							continue
						}
						pipeCountings[y+dir.Y][x+dir.X] = pipeCountings[y][x] + 1
						added = true
						if pipeCountings[y+dir.Y][x+dir.X] > highest {
							highest = pipeCountings[y+dir.Y][x+dir.X]
						}
					}
				}
			}
		}
		if !added {
			break
		}
	}

	// PrintMap(pipeMap)
	PrintMapNumbers(pipeCountings)

	println(highest)
}
