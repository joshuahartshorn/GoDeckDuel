package main

import "fmt"

// create a new type of deck which is a slice of strings
type deck [] string 

func newDeck() deck {
  cards := deck{}

  cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string{"Ace","Jack","Qeen", "King", "Two", "Three", "four", "Five", "Six", "Seven", "Eight", "Nine", "Ten" }

  for _, suit := range cardSuits {
	for _, value := range cardValues{
		cards = append (cards, value+" of "+ suit)
	}
  }
return cards
}

// create a new function that belongs to this deck type, whenever we call that function we should print out each indidual card in that deck
// A receiver 
 func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}

}