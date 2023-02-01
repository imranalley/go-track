package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	prodRate := float64(productionRate)
	workingCarsPerHour := (prodRate * successRate) 
	return (workingCarsPerHour/100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
	prodRate := float64(productionRate)
	workingCarsPerMinute := ((prodRate/60)*(successRate/100))
	return int(workingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
	const discountedCost = 95000
	const individualCost = 10000

	cost := (((carsCount/10)*discountedCost)+((carsCount%10)*individualCost))
	return uint(cost)
}
