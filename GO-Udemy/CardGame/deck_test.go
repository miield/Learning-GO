package main

import (
	"testing"
	"os"
)

func TestNewDeck(t *testing.T) {
	// Create a new deck
	d := newDeck()
	// check for the number of cards
	if len(d) != 16 {
		t.Errorf("Expected length of 16, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %v", d[0])
	}

	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected Four of Clubs but got %v", d[len(d)-1])
	}
}

// testing the saving of deck file to drive and retrieving it from the drive
func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	// remove any existing files called _decktesting
	os.Remove("_decktesting")

	// create a new deck
	deck := newDeck()
	// save it to the drive
	deck.saveToFile("_decktesting")

	// create a variable that loads the deck file from the drive
	loadDeckFile := newDeckFromFile("_decktesting")

	// check the number of cards in the loadDeckFile 
	if len(loadDeckFile) != 16 {
		t.Errorf("Expected 16, got %v", len(loadDeckFile))
	}

	os.Remove("_decktesting")
}
