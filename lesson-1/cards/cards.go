package main

import (
	"fmt"
	"math/rand"
	"time"
)
type Card struct {
	suit string
	value interface{}
}
type Deck []Card
func main(){
	deck:=GetCardDeck()
	deck.Shuffler()
	fmt.Println(deck)
}
func GetCardDeck()Deck{
	deck:=[]Card{}
	my_suits:=[]string{"Diamonds","Hearts","Spades","Clubs"}
	my_values:=[]string{"2","3","4","5","6","7","8","9","10","Jack","Queen","King","Ace"}
	for _, s:=range my_suits{
		for _, v:= range my_values{
			new_card:=Card {
				suit:s,
				value:v,
			}
			deck=append(deck,new_card)
		}
	}
	return deck
}
func (deck *Deck) Shuffler(){
	rand.Seed(time.Now().UnixNano()) //Set the seed

	rand.Shuffle(len(*deck), func(i, j int) { //call algorithm
		(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
	})

}
