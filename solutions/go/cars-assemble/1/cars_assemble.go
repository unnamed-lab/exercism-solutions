package cars

// import "fmt"

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	// panic("CalculateWorkingCarsPerHour not implemented")
	return float64(productionRate) * (successRate / 100.0)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// panic("CalculateWorkingCarsPerMinute not implemented")
	var result = float64(productionRate / 60) * (successRate / 100.0)

	return int(result)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	// panic("CalculateCost not implemented")
	
	const CostForGroup = 95000
	const CostForSingle = 10000
	
	var carsGroups int = carsCount / 10
	var carsSingle int = carsCount - (10 * carsGroups)
	
	// fmt.Println("Cars Count:: ", carsCount, "\nCost For Group:: ", CostForGroup, "\nCost For Single:: ", CostForSingle)

	var result = (carsGroups * CostForGroup) + (carsSingle * CostForSingle)

	return uint(result)
}
