package main

import (
	"os"
	"testing"
)

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

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	d := newDeck()
	d.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	deckLength := len(loadedDeck)

	if deckLength != 16 {
		t.Errorf("Expected deck length of 16, but got %v", deckLength)
	}

	os.Remove("_decktesting")
}
