package main

import (
	"fmt"
	"go_exercises/deck"
)

func main() {
	d := deck.New(deck.Deck(3))
	fmt.Println(d)
}
