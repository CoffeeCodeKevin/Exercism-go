package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, averagePrepTime int) int {
	var prepTime int
	if averagePrepTime == 0 {
		prepTime = len(layers) * 2
	} else {
		prepTime = len(layers) * averagePrepTime
	}
	return prepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0
	for i := 0; i < len(layers); i++ {
		switch layers[i] {
		case "noodles":
			noodles += 50
		case "sauce":
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList *[]string) {
	(*myList)[len(friendList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, portions int) []float64 {
	adjustedPortions := make([]float64, len(amounts))
	portionMult := float64(portions) / 2.0
	for i := 0; i < len(amounts); i++ {
		adjustedPortions[i] = amounts[i] * portionMult
	}
	return adjustedPortions
}
