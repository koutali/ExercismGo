package partyrobot

import (
	"fmt"
	"strconv"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return "Welcome to my party, " + name + "!"
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	ageString := strconv.Itoa(age)
	return "Happy birthday " + name + "! You are now " + ageString + " years old!"
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	// Welcome to my party, Christiane!
	// You have been assigned to table 027. Your table is on the left, exactly 23.8 meters from here.
	// You will be sitting next to Frank.

	welcomeString := Welcome(name)
	tableString := "\nYou have been assigned to table " + fmt.Sprintf("%03d", table) + "."
	secondTableString := "Your table is " + direction + ", exactly " + fmt.Sprintf("%.1f", distance) + " meters from here."
	neighborString := "\nYou will be sitting next to " + neighbor + "."

	return welcomeString +
		tableString +
		" " +
		secondTableString +
		neighborString
}
