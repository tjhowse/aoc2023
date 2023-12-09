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
	// main1()
	// main1(input)
	main2(input)
	// main2dumb(input)
	// end := time.Now()
	// println(end.Sub(start).Milliseconds())
}

func doMoves(location string, moves string, left map[string]string, right map[string]string) (int, string) {
	total := 0
	for _, dir := range moves {
		if dir == 'L' {
			location = left[location]
		} else {
			location = right[location]
		}
		// println(location)
		total += 1
	}
	return total, location
}
func doMovesChecking(location string, moves string, left map[string]string, right map[string]string) (int, string) {
	total := 0
	for _, dir := range moves {
		if dir == 'L' {
			location = left[location]
		} else {
			location = right[location]
		}
		// println(location)
		total += 1
		if location[2] == 'Z' {
			break
		}
	}
	return total, location
}

func checkLocations(locations []string) bool {
	for _, location := range locations {
		if location[2] != 'Z' {
			return false
		}
	}
	return true
}

func main2(inputfile string) {
	b := tj.FileToSlice(inputfile)

	// total := 0
	left := make(map[string]string)
	right := make(map[string]string)

	moves := b[0]
	locations := []string{}
	for _, line := range b[2:] {
		matches := tj.DoRegex(line, `(\w+) = \((\w+), (\w+)\)`)
		left[matches[0]] = matches[1]
		right[matches[0]] = matches[2]
		if matches[0][2] == 'A' {
			locations = append(locations, matches[0])
		}
	}
	println("locations", len(locations))
	steps := 0
	bigTotal := 1
	totals := []int{}
	wholePath := make(map[string]string)
	for _, location := range locations {
		total := 0
		startingLocation := location
		for {
			steps, location = doMovesChecking(location, moves, left, right)
			total += steps
			if location[2] == 'Z' {
				break
			}
		}
		wholePath[startingLocation] = location
		println("Starting location:", startingLocation, " end:", location, "took steps:", total)
		bigTotal *= total
		totals = append(totals, total)
	}

	// biggest := 0
	sumTotal := 1
	for _, total := range totals {
		// if bigTotal%total == 0 {
		// 	if total > biggest {
		// 		biggest = total
		// 	}
		// }
		println(total / len(moves))
		sumTotal *= total / len(moves)
	}

	println()
	println(sumTotal * len(moves))
	// println(bigTotal)
	// println(biggest)
	// 6890654361894480491 wrong
	// 6890654361894485488 wrong

}

func main2dumb(inputfile string) {
	b := tj.FileToSlice(inputfile)

	// total := 0
	left := make(map[string]string)
	right := make(map[string]string)

	moves := b[0]
	locations := []string{}
	for _, line := range b[2:] {
		matches := tj.DoRegex(line, `(\w+) = \((\w+), (\w+)\)`)
		left[matches[0]] = matches[1]
		right[matches[0]] = matches[2]
		if matches[0][2] == 'A' {
			locations = append(locations, matches[0])
		}
	}
	println("locations", len(locations))

	step := 0
	for {
		for i, location := range locations {
			if moves[step%len(moves)] == 'L' {
				locations[i] = left[location]
			} else {
				locations[i] = right[location]
			}
		}
		done := checkLocations(locations)
		// println()
		step += 1
		if step%1000000 == 0 {
			// if step%1 == 0 {
			println(step)
			for _, location := range locations {
				print(location, " ")
			}
			println()
		}
		if done {
			break
		}
	}
	println(step)
}

func main1(inputfile string) {
	b := tj.FileToSlice(inputfile)

	total := 0
	left := make(map[string]string)
	right := make(map[string]string)

	moves := b[0]
	for _, line := range b[2:] {
		matches := tj.DoRegex(line, `(\w+) = \((\w+), (\w+)\)`)
		left[matches[0]] = matches[1]
		right[matches[0]] = matches[2]
	}

	location := "AAA"
	steps := 0
	for {
		steps, location = doMoves(location, moves, left, right)
		total += steps
		if location == "ZZZ" {
			break
		}
	}

	println(total)
}
