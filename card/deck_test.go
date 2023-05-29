package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf(" Error lenth we are getting  %v", len(d))
	}

	if d[0] != "Spades  of  Ace" {
		t.Errorf(" We are getting %v", d[0])
	}

	if d[len(d)-1] != "clubs  of  King" {
		t.Errorf(" We are getting %v", d[len(d)-1])
	}
}

func TestSaveToFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 56 {
		t.Errorf("We are getting  %v", len(loadedDeck))
	}
}
