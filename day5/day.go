package main

import (
	"flag"
	"strings"
	"time"

	"github.com/tjhowse/tjgo"
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
	// main1()
	main1(input)
	// end := time.Now()
	// println(end.Sub(start).Milliseconds())
}

type translate struct {
	sourceRangeStart      int
	destinationRangeStart int
	rnge                  int
}

func (t *translate) FromString(s string) {
	split := strings.Split(s, " ")
	t.destinationRangeStart = tjgo.Str2intPanic(split[0])
	t.sourceRangeStart = tjgo.Str2intPanic(split[1])
	t.rnge = tjgo.Str2intPanic(split[2])
}

func (t *translate) InSourceRange(num int) bool {

	// t.Print()

	return num >= t.sourceRangeStart && num < t.sourceRangeStart+t.rnge
}
func (t *translate) Print() {
	println("sourceRangeStart", t.sourceRangeStart)
	println("destinationRangeStart", t.destinationRangeStart)
	println("rnge", t.rnge)
	println()
}

func (t *translate) Translate(num int) int {
	return t.destinationRangeStart + (num - t.sourceRangeStart)
}

func Translate(m []translate, source int) int {
	for _, t := range m {
		if t.InSourceRange(source) {
			return t.Translate(source)
		}
	}
	return source
}

func main1(inputfile string) {
	b := tj.FileToSlice(inputfile)

	maps := [][]translate{}
	for i := 0; i <= 7; i++ {
		maps = append(maps, []translate{})
	}

	stage := 0
	seedstrings := []string{}
	for _, line := range b {
		if line == "" {
			continue
		}
		split := strings.Split(line, " ")
		if len(split) == 0 {
			continue
		}
		if strings.HasPrefix(split[0], "seeds:") {
			stage = 0
		} else if strings.HasPrefix(split[0], "seed-to-soil") {
			stage = 1
			continue
		} else if strings.HasPrefix(split[0], "soil-to-fertilizer") {
			stage = 2
			continue
		} else if strings.HasPrefix(split[0], "fertilizer-to-water") {
			stage = 3
			continue
		} else if strings.HasPrefix(split[0], "water-to-light") {
			stage = 4
			continue
		} else if strings.HasPrefix(split[0], "light-to-temperature") {
			stage = 5
			continue
		} else if strings.HasPrefix(split[0], "temperature-to-humidity") {
			stage = 6
			continue
		} else if strings.HasPrefix(split[0], "humidity-to-location") {
			stage = 7
			continue
		}
		if stage == 0 {
			seedstrings = split[1:]
		} else {
			maps[stage-1] = append(maps[stage-1], translate{})
			maps[stage-1][len(maps[stage-1])-1].FromString(line)
		}

	}
	startTime := time.Now()
	// end := time.Now()
	// println(end.Sub(start).Milliseconds())
	// println(Translate(maps[0], 79))
	lowestLocation := 999999999999999
	// lowestLocationSeed := 0
	for j := 0; j < len(seedstrings); j += 2 {
		// println("Seed start: ", split[j])
		// println("Seed range: ", split[j+1])
		start := tj.Str2intPanic(seedstrings[j])
		rnge := tj.Str2intPanic(seedstrings[j+1])
		for seed := start; seed < start+rnge; seed++ {

			if time.Now().Sub(startTime).Seconds() > 600 {
				println("Timeout")
				break
			}
			// println("Seed", seed)
			next := seed
			for i := 0; i <= 7; i++ {
				next = Translate(maps[i], next)
				// println("Stage", i, next)
			}
			if next < lowestLocation {
				// lowestLocationSeed = seed
				lowestLocation = next
			}
		}
	}
	println(lowestLocation)
	// println(lowestLocationSeed)

	// 505472936 too high

}
