// Package weather provides tools to forecast the weather basic on the current location.
package weather

// CurrentCondition represents the current condition of the weather.
var CurrentCondition string

// CurrentLocation represents the current location of the user.
var CurrentLocation string

// Forecast accepts a city and a condition as string arguments and returns the current condition of the weather as string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
