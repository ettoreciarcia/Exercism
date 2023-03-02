// Package weather provides bla bla bla.
package weather

// CurrentCondition bla bla bla.
var CurrentCondition string

// CurrentLocation bla bla bla.
var CurrentLocation string

// Forecast returns bla bla bla.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
