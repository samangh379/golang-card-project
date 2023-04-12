package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)
type deck []string

func newDeck()deck{
	cards := deck{}

	cardSuits := deck{"Spades" , "Diamonds" , "Hearts" , "Club"}
    cardValues := deck {"One" , "Two" , "Three" , "Four"}

	for _ ,suits := range cardSuits{
		for _ ,values := range cardValues {

	cards = append(cards,values+" of "+suits)
		}
	}

	return cards
}

func (d deck) print(){
	for i , cards := range d {
		fmt.Println(i,cards)
	}
}

func  deal(d deck , handSize int) (deck ,deck){

	return d[:handSize] , d[handSize:]
}

func (d deck) toString() string {
	return strings.Join(d ,",")
}

func (d deck) saveToFiles(fileName string) error {
return ioutil.WriteFile(fileName ,[]byte(d.toString()), 0666)

}

func newDeckFromFile(fileName string) deck {
	bs, err :=ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}

	s := strings.Split(string(bs),",")

	return deck(s)

}

func (d deck)shuffle() {
	
	source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)
	
	for i := range d {
	newPosition := r.Intn(len(d) - 1)
	d[i],d[newPosition] = d[newPosition] , d[i]
	}
	
	}