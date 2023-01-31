package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
	actualMinutesInOven = (OvenTime - actualMinutesInOven)
	// panic("RemainingOvenTime not implemented")
	return actualMinutesInOven
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	// panic("PreparationTime not implemented")
	preparationTime := (numberOfLayers*2)
	return preparationTime
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	// panic("ElapsedTime not implemented")
	elapsedTime := (PreparationTime(numberOfLayers)+(40-RemainingOvenTime(actualMinutesInOven)))
	return elapsedTime
}
