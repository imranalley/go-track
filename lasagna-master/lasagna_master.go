package lasagna

// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, prepTime int) (totalPrepTime int) {
	if prepTime == 0 {
		prepTime = 2
	}
	totalPrepTime = (prepTime * len(layers))
	return
}

// TODO: define the 'Quantities()' function

func Quantities(layers []string) (noodles int, sauce float64) {
	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			noodles += 50
		} else if layers[i] == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function

func AddSecretIngredient(friendsList, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]

}

// TODO: define the 'ScaleRecipe()' function

func ScaleRecipe(quantities []float64, portions int) []float64 {
	var scaledQuantities []float64

	for _, quantity := range quantities {
		scaledQuantities = append(scaledQuantities, quantity/2.0*float64(portions))
	}
	return scaledQuantities
}
