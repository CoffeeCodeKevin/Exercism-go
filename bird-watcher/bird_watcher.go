package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	totalCount := 0
	for i := 0; i < len(birdsPerDay); i++ {
		totalCount += birdsPerDay[i]
	}
	return totalCount
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	totalCount := 0
	initialIndex := 0 + ((week - 1) * 7)
	for i := initialIndex; i < week*7; i++ {
		totalCount += birdsPerDay[i]
	}
	return totalCount
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	newSlice := birdsPerDay
	for i := 0; i < len(birdsPerDay); i++ {
		if i%2 == 0 {
			newSlice[i]++
		}
	}
	return newSlice
}
