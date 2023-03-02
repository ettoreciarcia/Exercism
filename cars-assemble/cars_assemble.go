package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	workingCars := float64(productionRate) * successRate / 100
	return workingCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	workingCars := (float64(productionRate) * successRate / 100) / 60
	return int(workingCars)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

	var cost uint

	switch {
	case carsCount < 10 && carsCount > 0:
		cost = uint(carsCount) * 10000
	case carsCount >= 10 && carsCount < 21:
		cost = uint(carsCount) * 9500
	case carsCount >= 21 && carsCount < 31:
		cost = uint(carsCount)*9000 + 11000
	case carsCount >= 31 && carsCount < 41:
		cost = uint(carsCount)*9000 + 22000
	case carsCount >= 41 && carsCount < 61:
		cost = uint(carsCount)*9000 + 31000
	case carsCount >= 100 && carsCount < 101:
		cost = uint(carsCount)*7500 + 200000
	case carsCount >= 148:
		cost = uint(carsCount)*7500 + 300000
	}

	return cost
}
