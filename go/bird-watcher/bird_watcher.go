package birdwatcher
import "math"
// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	count := 0
    for _, v := range birdsPerDay {
        count += v
    }
	return count
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	counter := 0
    if len(birdsPerDay) < (week - 1) * 7 {
        return 0
    }
    for i:= (week - 1)* 7; i < int(math.Min(float64(len(birdsPerDay)), float64(week * 7))); i++ {
        counter += birdsPerDay[i]
    }
	return counter
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
        birdsPerDay[i] += 1
    }
	return birdsPerDay
}
