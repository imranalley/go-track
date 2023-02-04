package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
	switch card {
	case "ace":
		return 11
	case "king":
		return 10
	case "queen":
		return 10
	case "jack":
		return 10
	case "ten":
		return 10
	case "nine":
		return 9
	case "eight":
		return 8
	case "seven":
		return 7
	case "six":
		return 6
	case "five":
		return 5
	case "four":
		return 4
	case "three":
		return 3
	case "two":
		return 2
	}
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")
	Sum := ParseCard(card1) + ParseCard(card2)
	// fmt.Sprintf("Sum is = %d", Sum)
	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case (Sum == 21) && (ParseCard(dealerCard) == 10 || ParseCard(dealerCard) == 11):
		return "S"
	case (Sum == 21) && (ParseCard(dealerCard) != 10 || ParseCard(dealerCard) != 11):
		return "W"
	case Sum == 17 || Sum == 18 || Sum == 19 || Sum == 20:
		return "S"
	case (Sum == 12 || Sum == 13 || Sum == 14 || Sum == 15 || Sum == 16) && ParseCard(dealerCard) >= 7:
		return "H"
	case Sum == 12 || Sum == 13 || Sum == 14 || Sum == 15 || Sum == 16:
		return "S"
	case Sum <= 11:
		return "H"
	}
	return ""
}
