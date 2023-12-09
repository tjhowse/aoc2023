package main

import (
	"flag"
	"time"
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

func main1(inputfile string) {
	// b := tj.FileToSlice(inputfile)

	// raceTimes := []int{7, 15, 30}
	// dist := []int{9, 40, 200}
	// raceTimes := []int{42, 89, 91, 89}
	// dist := []int{308, 1170, 1291, 1467}
	// raceTimes := []int{71530}
	// dist := []int{940200}
	raceTimes := []int{42899189}
	dist := []int{308117012911467}

	startTime := time.Now()
	total := 1
	for i, raceTime := range raceTimes {
		ways := 0
		for holdTime := 0; holdTime <= raceTime; holdTime++ {
			d := holdTime * (raceTime - holdTime)
			if d > dist[i] {
				ways++
			}
			if time.Now().Sub(startTime).Seconds() > 600 {
				println("Timeout")
				break
			}
		}
		total *= ways
	}
	println(total)
}
