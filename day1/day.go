package main

import (
	"strconv"
	"strings"

	tj "github.com/tjhowse/tjgo"
)

func main() {
	main2()
}

func replace_substring(line string) (bool, string) {
	lookup := make(map[string]string)
	lookup["one"] = "1"
	lookup["two"] = "2"
	lookup["three"] = "3"
	lookup["four"] = "4"
	lookup["five"] = "5"
	lookup["six"] = "6"
	lookup["seven"] = "7"
	lookup["eight"] = "8"
	lookup["nine"] = "9"

	result := line
	for i := 0; i < len(line); i++ {
		for key, value := range lookup {
			if i+len(key) > len(line) {
				continue
			}
			if line[i:i+len(key)] == key {
				result = strings.Replace(result, key, value, -1)
				return true, result
			}
		}
	}
	return false, result
}

func main2bad() {
	b := tj.FileToSlice("input_real")
	// b := tj.FileToSlice("input")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DIntSlice("input", ' ')

	first := ""
	last := ""
	total := 0

	filtered := make([]string, 0)

	for _, line := range b {
		println(line)
		result := line
		replaced := false
		for {
			replaced, result = replace_substring(result)
			if !replaced {
				break
			}
		}
		println(result)
		filtered = append(filtered, result)
	}

	for _, line := range filtered {
		// println(line)
		first = ""
		last = ""
		for _, rune := range line {
			_, err := strconv.ParseInt(string(rune), 10, 64)
			if err != nil {
				continue
			}
			if first == "" {
				first = string(rune)
			}
			last = string(rune)
		}
		// println(line)
		// println(first + last)
		total += tj.Str2int(first + last)
	}
	// print(total)

	// 53306 wrong
	// 53337 wrong
	// 52572 wrong
}

func main2() {
	lookup := make(map[string]string)
	lookup["one"] = "1"
	lookup["two"] = "2"
	lookup["three"] = "3"
	lookup["four"] = "4"
	lookup["five"] = "5"
	lookup["six"] = "6"
	lookup["seven"] = "7"
	lookup["eight"] = "8"
	lookup["nine"] = "9"
	b := tj.FileToSlice("input_real")
	// b := tj.FileToSlice("input")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DIntSlice("input", ' ')

	first := ""
	last := ""
	total := 0

	for _, line := range b {
		first = ""
		last = ""
		for i, rune := range line {
			_, err := strconv.ParseInt(string(rune), 10, 64)
			if err != nil {
				// found := false
				for key, value := range lookup {
					if i+len(key) > len(line) {
						continue
					}
					if line[i:i+len(key)] == key {
						// found = true
						// println("Found a match at", i, key, value)
						if first == "" {
							first = value
						}
						last = value
						break
					}
				}
				continue
			}
			if first == "" {
				first = string(rune)
			}
			last = string(rune)
		}
		// println(first + last)
		total += tj.Str2int(first + last)
	}
	print(total)

}

func main1() {
	b := tj.FileToSlice("input_real")
	// b := tj.FileToSlice("input")
	// b := tj.FileToIntSlice("input")
	// b := tj.FileTo2DIntSlice("input", ' ')

	first := ""
	last := ""
	total := 0

	for _, line := range b {
		first = ""
		last = ""
		for _, rune := range line {
			_, err := strconv.ParseInt(string(rune), 10, 64)
			if err != nil {
				continue
			}
			if first == "" {
				first = string(rune)
			}
			last = string(rune)
		}
		// println(first + last)
		total += tj.Str2int(first + last)
	}
	print(total)

}
