package techpalace

import (
	"fmt"
	"strings"
)

const WELCOME_MESSAGE = "Welcome to the Tech Palace,"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("%s %s", WELCOME_MESSAGE, strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	s := ""

	for numStarsPerLine > 0 {
		s += "*"
		numStarsPerLine -= 1
	}

	return fmt.Sprintf("%s\n%s\n%s", s, welcomeMsg, s)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// res := " "

	// for i := 0; i < len(oldMsg); i++ {
	// 	if oldMsg[i] == '*' || oldMsg[i] == '\n' {
	// 		continue
	// 	}

	// 	res += string(oldMsg[i])
	// }

	return strings.Trim(oldMsg, " *\n")
}
