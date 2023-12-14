package main

import (
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	// _ is a placeholder for variables that we don't care about
	// in this case, we don't care about the index
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	// Slices are zero-indexed
	// [0:handSize] means from 0 to handSize, not including handSize
	// [handSize:] means from handSize to the end of the slice
	return d[:handSize], d[handSize:]
}

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return s
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
