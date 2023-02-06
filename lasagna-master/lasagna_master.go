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

func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	scaledQuantities = quantities[:]
	for i := 0; i < len(scaledQuantities); i++ {
		scaledQuantities[i] = scaledQuantities[i] * (float64(portions) / 2.0)
	}
	return
}
