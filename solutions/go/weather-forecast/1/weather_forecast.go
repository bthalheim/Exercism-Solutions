// Package weather test weather package.
package weather

// CurrentCondition defines the current weather condition. 
var CurrentCondition string

// CurrentLocation defines the current geo location.
var CurrentLocation string

// Forecast returns the current forcast for a given location. It takes the city and current condition as arguments.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
