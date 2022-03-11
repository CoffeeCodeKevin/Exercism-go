package scrabble

import "strings"

func letterScore(letter string) int {
	pointDict := make(map[string]int)
	pointDict["A"] = 1
	pointDict["E"] = 1
	pointDict["I"] = 1
	pointDict["O"] = 1
	pointDict["U"] = 1
	pointDict["L"] = 1
	pointDict["N"] = 1
	pointDict["R"] = 1
	pointDict["S"] = 1
	pointDict["T"] = 1
	pointDict["D"] = 2
	pointDict["G"] = 2
	pointDict["B"] = 3
	pointDict["C"] = 3
	pointDict["M"] = 3
	pointDict["P"] = 3
	pointDict["F"] = 4
	pointDict["H"] = 4
	pointDict["V"] = 4
	pointDict["W"] = 4
	pointDict["Y"] = 4
	pointDict["K"] = 5
	pointDict["J"] = 8
	pointDict["X"] = 8
	pointDict["Q"] = 10
	pointDict["Z"] = 10
	point := pointDict[letter]
	return point
}

func Score(word string) int {
	totalCount := 0
	for i := 0; i < len(word); i++ {
		totalCount += letterScore(strings.ToUpper(string(word[i])))
	}
	return totalCount
}
