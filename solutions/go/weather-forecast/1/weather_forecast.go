// package for printing wether forecasts
package weather

var (
	// Contains certain state of wether
	CurrentCondition string
	// Contain location's name what will be used for forecast
	CurrentLocation string
)

// Prints connected information about forecast, where it is, and what it is
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
