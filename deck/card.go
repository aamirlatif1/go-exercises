//go:generate stringer -type=Suit,Rank
package deck

import "fmt"

// Suit - custom type to hold value for card suit (Spade, Heart, Diamond and Club)
type Suit int8
type Rank int8

// Declares related constants for each suit starting with index 1
const (
	Spade Suit = iota
	Diamond
	Club
	Heart
	Joker // Jocker is a spacial case
)

var suits = [...]Suit{Spade, Diamond, Club, Heart}

const (
	_ Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	minRank = Ace
	maxRank = King
)

// Card is a combination of suit and rank
type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %s", c.Rank.String(), c.Suit.String())
}
