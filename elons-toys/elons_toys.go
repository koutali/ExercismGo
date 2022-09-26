package elon

import "strconv"

// TODO: define the 'Drive()' method
func (car *Car) Drive() {
	if car == nil {
		return
	}

	if car.battery >= car.batteryDrain {
		car.distance += car.speed
		car.battery -= car.batteryDrain
	}
}

// TODO: define the 'DisplayDistance() string' method
func (car *Car) DisplayDistance() string {
	if car == nil {
		panic("car cannot be nil")
	}

	distanceMessage := "Driven " + strconv.Itoa(car.distance) + " meters"
	return distanceMessage
}

// TODO: define the 'DisplayBattery() string' method
func (car *Car) DisplayBattery() string {
	if car == nil {
		panic("car cannot be nil")
	}

	distanceMessage := "Battery at " + strconv.Itoa(car.battery) + "%"
	return distanceMessage
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
// CanFinish checks if a car is able to finish a certain track.
func (car *Car) CanFinish(trackDistance int) bool {
	if car == nil {
		panic("car cannot be nil")
	}

	if car.speed == 0 {
		return false
	}

	time := trackDistance / car.speed
	leftBattery := car.battery - (time * car.batteryDrain)
	return leftBattery >= 0
}
