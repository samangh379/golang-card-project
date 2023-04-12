package main

import (
	"os"
	"testing"
)


func TestNewDeck(t *testing.T) {
	d :=newDeck()

	if len(d) != 16 {
		t.Errorf("we expect got 16 card but we got %v", len(d))
	}

}


func TestSaveToFileAndReadFile(t *testing.T){
	os.Remove("_deckTesting")
deck :=newDeck()
deck.saveToFiles("_deckTesting")
loadedDeck :=newDeckFromFile("_deckTesting")
if len(loadedDeck) != 16 {
	t.Errorf("we expect got 16 cards but we got %v",len(loadedDeck))
}
}