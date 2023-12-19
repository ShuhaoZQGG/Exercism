// Package weather provides tools to forecase weather conditon in cities.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string

// CurrentLocation represents the location of the current city.
var CurrentLocation string

// Forecast takes two arguments, string city and condition, returns a forecast of current condition and the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
