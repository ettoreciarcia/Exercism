package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	characters := "*"
	return strings.Repeat(characters, numStarsPerLine) + "\n" + welcomeMsg + "\n" + strings.Repeat(characters, numStarsPerLine)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.Trim(oldMsg, "\n*"))

}
