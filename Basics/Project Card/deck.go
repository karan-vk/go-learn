package main

import (
	"math/rand"
	"time"
	// "fmt"
	"io/ioutil"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "King", "Queen"}
	for _, suit := range cardSuits {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suit)

		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		println(i+1, card)
		
	}
}
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]

}
func (d deck) toString() string {
	return strings.Join([]string(d), ",")

}
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		println(err)
		os.Exit(1)
		
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}

func (d deck) shuffle() {
	source:=rand.NewSource(time.Now().UnixNano())
	r:=rand.New(source)
	
	for i:=range d{
		newPosition:=r.Intn(len(d)-1)
		
		
		d[i],d[newPosition]=d[newPosition],d[i]

	}

}
