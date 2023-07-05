// Package weather forecast the current weather in various cities of Goblinocus.
package weather

// CurrentCondition represents the current weather.
var CurrentCondition string

// CurrentLocation represents the current city location.
var CurrentLocation string

// Forecast returns the current weather in the considered location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
