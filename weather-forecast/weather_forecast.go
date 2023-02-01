// Package weather allows you to pass a City's name and condition and print out a forecast.
package weather
// CurrentCondition is a public string variable that contains the current weather condition.
var CurrentCondition string
// CurrentLocation is a public string variable that containes the current location.
var CurrentLocation string
// Forecast returns a string value that displays the current weather conditions and location that was passed through.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
