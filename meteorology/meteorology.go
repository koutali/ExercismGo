package meteorology

import "fmt"

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (tempUnit TemperatureUnit) String() string {
	units := []string{"°C", "°F"}
	return units[tempUnit]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

func (temperature Temperature) String() string {
	return fmt.Sprintf("%d%s%s", temperature.degree, " ", temperature.unit.String())
}

// Add a String method to the Temperature type
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedUnit SpeedUnit) String() string {
	units := []string{"km/h", "mph"}
	return units[speedUnit]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
	return fmt.Sprintf("%d%s%s", speed.magnitude, " ", speed.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

func (meteoData MeteorologyData) String() string {
	location := fmt.Sprintf("%s%s", meteoData.location, ": ")
	tempData := fmt.Sprintf("%s%s", meteoData.temperature.String(), ", ")
	windData := fmt.Sprintf("%s%s%s%s%s", "Wind ", meteoData.windDirection, " at ", meteoData.windSpeed.String(), ", ")
	humidityData := fmt.Sprintf("%d%s", meteoData.humidity, "% Humidity")

	return location + tempData + windData + humidityData
}

// Add a String method to MeteorologyData
