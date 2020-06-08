package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T)  {
	d:=newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v",len(d))

	}
	if d[0] != "Ace of Spades"{
		t.Errorf("Expected first card Ace of spades but got %v",d[0])
	}
	if d[len(d)-1] != "Queen of Clubs"{
		t.Errorf("Expected last card of Four of Clubs but got %v",d[len(d)])
	}
	
	
}
func TestSaveToDeckAndNewDeckFromFile(t *testing.T){
	os.Remove("_dectesting")
	deck:=newDeck()
	deck.saveToFile("_dectesting")

	loadedDeck:=newDeckFromFile("_dectesting")
	if len(loadedDeck)!= 52{
		t.Error("error at save")
	}
	os.Remove("_dectesting")


}