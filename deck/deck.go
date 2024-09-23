package deck

import (
	"sort"
	"time"

	"golang.org/x/exp/rand"
)

func New(opts ...func([]Card) []Card) []Card {
	var cards []Card
	for _, suit := range suits {
		for r := minRank; r <= maxRank; r++ {
			cards = append(cards, Card{
				Rank: Rank(r),
				Suit: Suit(suit),
			})
		}
	}
	for _, opt := range opts {
		cards = opt(cards)
	}
	return cards
}

func Filter(f func(card Card) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for _, c := range cards {
			if !f(c) {
				ret = append(ret, c)
			}
		}
		return ret
	}
}

func Suffle(cards []Card) []Card {
	ret := make([]Card, len(cards))
	r := rand.New(rand.NewSource(uint64(time.Now().Unix())))
	perms := r.Perm(len(cards))
	for i, p := range perms {
		ret[i] = cards[p]
	}
	return ret
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	})
	return cards
}

func absRank(c Card) int {
	return int(c.Suit)*int(maxRank) + int(c.Rank)
}

func Sort(less func(cards []Card) func(i, j int) bool) func([]Card) []Card {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

func Jockers(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}

func Deck(n int) func([]Card) []Card {
	return func(cards []Card) []Card {
		var ret []Card
		for i := 0; i < n; i++ {
			ret = append(ret, cards...)
		}
		return ret
	}
}
