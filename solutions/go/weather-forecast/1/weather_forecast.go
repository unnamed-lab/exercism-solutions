// Package weather provides a function to forecast the weather condition of a city.
package weather

var (
	// CurrentCondition represents the current weather condition of the city.
	CurrentCondition string
	// CurrentLocation represents the current location of the city.
	CurrentLocation  string
)

// Forecast returns the current weather condition of the city.
// It takes the city and the condition as input and returns the current weather condition of the city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
