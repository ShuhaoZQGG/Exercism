package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var output float64
    output = float64(productionRate) * successRate / 100.0
    return output
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	costOfGroupCars := 95000
    const costOfIndividualCars = 10000
    var numberOfGroupCars int = carsCount / 10
    var numberOfIndividualCars int = carsCount % 10
    var output uint = uint(costOfGroupCars * numberOfGroupCars + numberOfIndividualCars * costOfIndividualCars)
    return output
}
