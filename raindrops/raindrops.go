package raindrops

import "strconv"

func Convert(number int) string {
	raindrops := ""
	factorial := false
	if number%3 == 0 {
		raindrops += "Pling"
		factorial = true
	}
	if number%5 == 0 {
		raindrops += "Plang"
		factorial = true
	}
	if number%7 == 0 {
		raindrops += "Plong"
		factorial = true
	}
	if !factorial {
		raindrops = strconv.Itoa(number)
	}

	return raindrops
}
