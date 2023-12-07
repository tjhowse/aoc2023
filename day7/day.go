package main

import (
	"flag"
	"sort"
	"strings"

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

type hand struct {
	s      string
	cards  [5]rune
	binned map[rune]int
	bid    int
}

func (h *hand) FromStr(s string) {
	h.s = s
	h.binned = make(map[rune]int)
	spl := strings.Split(s, " ")
	h.bid = tj.Str2int(spl[1])
	for i, c := range spl[0] {
		h.cards[i] = c
	}
	h.BinCards()
	// println(s)
}

func CountInHand(h [5]rune, c rune) int {
	count := 0
	for _, card := range h {
		if card == c {
			count++
		}
	}
	return count
}

func (h *hand) BinCards() {
	for _, c := range h.cards {
		h.binned[c] = CountInHand(h.cards, c)
	}
}

func (h *hand) Calc() int {

	if h.binned['J'] == 0 {
		return CalcImpl(h.binned)
	}

	// Replace J with the most common other card
	max := 0
	maxCard := ' '
	for card, count := range h.binned {
		if count > max && card != 'J' {
			max = count
			maxCard = card
		}
	}
	if max == 1 {

	}
	// println("Rebinning J to", string(maxCard), "in", h.s)
	rebinned := h.binned
	rebinned[maxCard] += h.binned['J']
	rebinned['J'] = 0
	// for k, v := range rebinned {
	// 	println(string(k), v)
	// }
	return CalcImpl(rebinned)

}

func CalcImpl(binned map[rune]int) int {
	/*
		6: 5 of kind
		5: 4 of kind
		4: full house
		3: 3 of kind
		2: 2 pair
		1: 1 pair
		0: no dupes
	*/

	counts := make(map[int]int)

	for _, count := range binned {
		counts[count] += 1
	}

	if counts[5] > 0 {
		return 6
	}
	if counts[4] > 0 {
		return 5
	}
	if counts[3] > 0 && counts[2] > 0 {
		return 4
	}
	if counts[3] > 0 {
		return 3
	}
	if counts[2] == 2 {
		return 2
	}
	if counts[2] == 1 {
		return 1
	}

	return 0
}

func RuneToScore(r rune) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		// return 11
		return 1
	case 'T':
		return 10
	default:
		return tj.Str2int(string(r))
	}
}

func Compare(h1 *hand, h2 *hand) int {
	if h1.Calc() > h2.Calc() {
		return 1
	}
	if h1.Calc() < h2.Calc() {
		return -1
	}
	for i := 0; i < len(h1.cards); i++ {
		if RuneToScore(h1.cards[i]) > RuneToScore(h2.cards[i]) {
			return 1
		}
		if RuneToScore(h1.cards[i]) < RuneToScore(h2.cards[i]) {
			return -1
		}
	}

	return 0
}

func main1(inputfile string) {
	hands := []hand{}
	b := tj.FileToSlice(inputfile)
	total := 0
	for _, line := range b {
		hands = append(hands, hand{})
		hands[len(hands)-1].FromStr(line)
		// println(hands[len(hands)-1].Calc())
	}

	sort.Slice(hands, func(i, j int) bool { return Compare(&hands[i], &hands[j]) > 0 })

	for i, h := range hands {
		rank := len(hands) - i
		// println("rank", len(hands)-i, h.s)
		// println(h.s)
		total += rank * h.bid
	}

	println(total)
	// 253181716 wrong
	// 253473930 right
}
