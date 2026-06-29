// Package weather contains variables and a function related to the forecast.
package weather

var (
    // CurrentCondition holds a string value representing the current weather condition.
	CurrentCondition string
    // CurrentLocation holds a string value and represents the current location.
	CurrentLocation  string
)

// Forecast takes a city and a condition as arguments, and returns the current weather condition of
// that specific city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
