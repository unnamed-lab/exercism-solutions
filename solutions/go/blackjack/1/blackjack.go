package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
	switch card {
		case "ace":
			return 11
		case "two":
			return 2
		case "three":
			return 3
		case "four":
			return 4
		case "five":
			return 5
		case "six":
			return 6
		case "seven":
			return 7
		case "eight":
			return 8
		case "nine":
			return 9
		case "ten":
			return 10
		case "jack":
			return 10
		case "queen":
			return 10
		case "king":
			return 10
		default:
			return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")

	/*
    If you have a pair of aces you must always split them.
    If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, 
		a face card (Jack/Queen/King) or a ten then you automatically win. 
	If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
    If your cards sum up to a value within the range [17, 20] you should always stand.
    If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, 
		in which case you should always hit.
    If your cards sum up to 11 or lower you should always hit.
	*/

	playerScore := ParseCard(card1) + ParseCard(card2)
	dealerScore := ParseCard(dealerCard)

	switch {
	case playerScore == 22:
		return "P"
	case playerScore == 21 && dealerScore < 10:
		return "W"
	case playerScore >= 17 && playerScore <= 20:
		return "S"
	case playerScore >= 12 && playerScore <= 16 && dealerScore >= 7:
		return "H"
	case playerScore <= 11:
		return "H"
	default:
		return "S"
	}


}
