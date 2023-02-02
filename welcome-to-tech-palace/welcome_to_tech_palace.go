package techpalace
import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	// panic("Please implement the WelcomeMessage() function")
	return ("Welcome to the Tech Palace, " + strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	// panic("Please implement the AddBorder() function")
	stars := strings.Repeat("*", numStarsPerLine)
	welcome := (stars + "\n" + welcomeMsg + "\n" + stars)
	return(welcome)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	// panic("Please implement the CleanupMessage() function")
	newMsg := strings.Trim(oldMsg, "*")
	newMsgSpace := strings.TrimSpace(newMsg)
	newMsgSpace = strings.TrimPrefix(newMsgSpace, "*")
	newMsgSpace = strings.TrimSuffix(newMsgSpace, "*")
	newMsgSpace = strings.TrimSpace(newMsgSpace)
	return(newMsgSpace)
}
