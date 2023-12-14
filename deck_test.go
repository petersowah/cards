package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	deckLength := len(d)

	if deckLength != 16 {
		t.Errorf("Expected deck length of 16, but got %v", deckLength)
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[deckLength-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[deckLength-1])
	}
}
