package main

import (
	"strings"
	"time"

	tj "github.com/tjhowse/tjgo"
)

func main() {

	start := time.Now()
	main1()
	end := time.Now()
	println(end.Sub(start).Milliseconds())
}

type game struct {
	id    int
	red   int
	green int
	blue  int
}

func (g *game) print() {
	println("id", g.id, "red", g.red, "green", g.green, "blue", g.blue)
}

func (g *game) power() int {
	return g.red * g.green * g.blue
}

func (g *game) FromString(s string) {
	c := tj.DoRegex(s, `Game (\d+): (.*)`)
	// println(c[0])
	g.id = tj.Str2int(c[0])
	round := strings.Split(c[1], "; ")
	for _, i := range round {
		red := tj.DoRegex(i, `(\d+) red`)
		if len(red) > 0 {
			// println(red[0])
			r := tj.Str2int(red[0])
			if r > g.red {
				g.red = r
			}
		}
		green := tj.DoRegex(i, `(\d+) green`)
		if len(green) > 0 {
			// println(green[0])
			r := tj.Str2int(green[0])
			if r > g.green {
				g.green = r
			}
		}
		blue := tj.DoRegex(i, `(\d+) blue`)
		if len(blue) > 0 {
			// println(blue[0])
			r := tj.Str2int(blue[0])
			if r > g.blue {
				g.blue = r
			}
		}
	}
}

func main1() {
	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DIntSlice("input", ' ')

	total := 0

	for _, line := range b {
		g := game{}
		g.FromString(line)
		// g.print()
		// if !(g.red > 12 || g.green > 13 || g.blue > 14) {
		// 	total += g.id
		// 	// println(g.id, "possible")
		// }
		// println(g.id, g.power())
		total += g.power()
	}
	println(total)

}
