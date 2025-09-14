package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return successRate * (float64(productionRate) / 100.00) 
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int((successRate / 100.00) * (float64(productionRate) / 60.00))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {

    // 10000 Per Car 
    // 95000 For A Batch of 10 cars  

    var grossCost int = carsCount * 10000
    var batchDiscount int = 5000 * (carsCount / 10) 
    var netCost = uint(grossCost - batchDiscount)

    return netCost
}
