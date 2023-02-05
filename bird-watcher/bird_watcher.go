package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	// panic("Please implement the TotalBirdCount() function")
	totalBirds := 0
	for i := 0; i < cap(birdsPerDay); i++ {
		totalBirds += birdsPerDay[i]
	}
	return totalBirds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function")
	daysTracked := week * 7
	totalBirds := 0
	for i := 0; i < 7; i++ {
		totalBirds += birdsPerDay[daysTracked-1]
		daysTracked -= 1
	}
	return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	// panic("Please implement the FixBirdCountLog() function")
	for i := 0; i < cap(birdsPerDay); i = i + 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay[:]
}
