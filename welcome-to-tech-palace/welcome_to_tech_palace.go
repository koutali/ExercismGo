package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	if numStarsPerLine < 0 {
		panic("numStarsPerLine is negative")
	}

	var message string = strings.Repeat("*", numStarsPerLine) +
		"\n" +
		welcomeMsg + "\n" +
		strings.Repeat("*", numStarsPerLine)

	return message

}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {

	if strings.EqualFold(oldMsg, "") {
		panic("oldMsg is empty")
	}

	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
	oldMsg = strings.TrimSpace(oldMsg)

	return oldMsg
}
