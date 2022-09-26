package purchase

import "strings"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {

	var lowerCase = strings.ToLower(kind)
	return strings.EqualFold(lowerCase, "truck") || strings.EqualFold(lowerCase, "car")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {

	var selectedCar string
	if option1 < option2 {
		selectedCar = option1
	} else {
		selectedCar = option2
	}

	selectedCar += " is clearly the better choice."
	return selectedCar
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.8
	} else if age >= 3 && age < 10 {
		return originalPrice * 0.7
	}

	return originalPrice * 0.5
}
