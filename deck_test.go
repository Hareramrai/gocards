package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	cards := newDeck()

	if len(cards) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(cards))
	}

	if cards[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades, but got %v", cards[0])
	}

	if cards[len(cards)-1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs, but got %v", cards[len(cards)-1])
	}
}

func TestDeal(t *testing.T) {

	cards := newDeck()

	d1, d2 := deal(cards, 5)

	if len(d1) != 5 {
		t.Errorf("Expected 5 cards, but got %v", len(d1))
	}

	if len(d2) != len(cards)-5 {
		t.Errorf("Expected %d cards, but got %v", len(cards)-5, len(d2))
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
