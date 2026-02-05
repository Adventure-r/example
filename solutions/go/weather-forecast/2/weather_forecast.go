// Package weather for printing wether forecasts.
package weather

var (
	// CurrentCondition contains certain state of wether.
	CurrentCondition string
	// CurrentLocation contain location's name what will be used for forecast.
	CurrentLocation string
)

// Forecast prints connected information about forecast, where it is, and what it is.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
