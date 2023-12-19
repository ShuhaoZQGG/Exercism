package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
    		return 11
        case "king", "queen", "jack", "ten":
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
        default:
    		return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var sum int = ParseCard(card1) + ParseCard(card2)
    var dealerValue int = ParseCard(dealerCard)
    switch {
        case sum >= 17 && sum <= 20:
    		return "S"
        case sum <= 11:
    		return "H"
        case sum == 21:
    		if dealerValue < 10 {
                return "W"
            }
    		return "S"
        case sum >= 12 && sum <= 17:
			if dealerValue >= 7 {
                return "H"
            }
    		return "S"
        default:
    		return "P"
    }
}
