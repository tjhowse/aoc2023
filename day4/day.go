package main

import (
	"flag"
	"strings"

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
	main2(input)
	// end := time.Now()
	// println(end.Sub(start).Milliseconds())
}

func get_score2(winner []int, haver []int) int {
	score := 0
	for _, w := range winner {
		for _, h := range haver {
			if h != w {
				continue
			}
			score += 1
		}
	}
	return score
}

func recurse_scores(scores map[int]int, game_number int) int {
	total := 0
	score := scores[game_number]
	total += 1
	// print("Game ", game_number, " wins cards: ")
	// for i := game_number + 1; i <= game_number+score; i++ {
	// 	print(i, " ")
	// }
	// println()
	for i := game_number + 1; i <= game_number+score; i++ {
		total += recurse_scores(scores, i)
	}
	// println("Game", game_number, "Total:", total)
	// println()
	return total
}

func main2(inputfile string) {
	b := tj.FileToSlice(inputfile)

	scores := map[int]int{}
	total := 0
	max_game := 0
	// count := 0
	for _, line := range b {
		// println(line)
		winners := []int{}
		havers := []int{}
		game_number := tjgo.Str2int(tjgo.DoRegex(line, `^Card[ ]+(\d+):`)[0])
		winner_string := tjgo.DoRegex(line, `.*:([ \d]*) \|`)[0]
		for _, num := range strings.Split(winner_string, " ") {
			if len(num) == 0 {
				continue
			}
			winners = append(winners, tjgo.Str2int(num))
		}
		haver_string := tjgo.DoRegex(line, `\| ([ \d]*)$`)
		for _, num := range strings.Split(haver_string[0], " ") {
			if len(num) == 0 {
				continue
			}
			havers = append(havers, tjgo.Str2int(num))
		}
		// println(game_number, winners, havers)
		score := get_score2(winners, havers)
		// println("game number", game_number, ":", score)
		scores[game_number] = score
		if game_number > max_game {
			max_game = game_number
		}
	}
	for i := 1; i <= max_game; i++ {
		total += recurse_scores(scores, i)
	}
	println(total)

}

func get_score(winner []int, haver []int) int {
	score := 0
	for _, w := range winner {
		for _, h := range haver {
			if h != w {
				continue
			}
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	return score
}

func main1() {
	// b := tj.FileToSlice("input")
	b := tj.FileToSlice("input_real")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DIntSlice("input", ' ')

	scores := map[int]int{}
	total := 0
	// count := 0
	for _, line := range b {
		println(line)
		winners := []int{}
		havers := []int{}
		game_number := tjgo.Str2int(tjgo.DoRegex(line, `^Card[ ]+(\d+):`)[0])
		winner_string := tjgo.DoRegex(line, `.*:([ \d]*) \|`)[0]
		for _, num := range strings.Split(winner_string, " ") {
			if len(num) == 0 {
				continue
			}
			winners = append(winners, tjgo.Str2int(num))
		}
		haver_string := tjgo.DoRegex(line, `\| ([ \d]*)$`)
		for _, num := range strings.Split(haver_string[0], " ") {
			if len(num) == 0 {
				continue
			}
			havers = append(havers, tjgo.Str2int(num))
		}
		// println(game_number, winners, havers)
		score := get_score(winners, havers)
		// println("game number", game_number, ":", score)
		total += score
		scores[game_number] = score
	}
	println(total)

}
