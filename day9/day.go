package main

import (
	"flag"

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

func PrintSequence(in []int) {
	for _, i := range in {
		print(i, " ")
	}
	println()
}

func GetSequenceDifferences(in []int) []int {
	out := []int{}
	for i := 1; i < len(in); i++ {
		out = append(out, in[i]-in[i-1])
	}
	return out
}

func AllZeroes(in []int) bool {
	for _, i := range in {
		if i != 0 {
			return false
		}
	}
	return true
}

func CalculateDifferencesUntilAllZero(in []int) int {
	next := GetSequenceDifferences(in)
	if AllZeroes(next) {
		return in[0]
	}
	// PrintSequence(in)
	return in[len(in)-1] + CalculateDifferencesUntilAllZero(next)
}
func CalculateDifferencesUntilAllZeroPart2(in []int) int {
	next := GetSequenceDifferences(in)
	if AllZeroes(next) {
		return in[0]
	}
	// PrintSequence(in)
	return in[0] - CalculateDifferencesUntilAllZeroPart2(next)
}

func main2(inputfile string) {
	b := tj.FileTo2DIntSliceWhole(inputfile, ' ')

	total := 0

	for _, line := range b {
		result := CalculateDifferencesUntilAllZeroPart2(line)
		total += result
	}

	println(total)
}
func main1(inputfile string) {
	b := tj.FileTo2DIntSliceWhole(inputfile, ' ')

	total := 0

	for _, line := range b {
		result := CalculateDifferencesUntilAllZero(line)
		total += result
	}

	println(total)
}
