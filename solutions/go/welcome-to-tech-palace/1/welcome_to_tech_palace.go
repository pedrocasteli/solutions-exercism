package techpalace

import "strings"


// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	msg := ""

	for i := 0; i < numStarsPerLine; i++ {
		msg += "*"
	}

	msg += "\n" + welcomeMsg + "\n"

	for i := 0; i < numStarsPerLine; i++ {
		msg += "*"
	}

	return msg
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	msgNoStars := strings.ReplaceAll(oldMsg, "*", "")
	msgNoSpaces := strings.TrimSpace(msgNoStars)

	return msgNoSpaces
}
