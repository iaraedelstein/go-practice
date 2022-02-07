package main

import "fmt"

func main(){
	cards := newDeck()

	// hand, remainingDeck := deal(cards, 5)
	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())

	//cards.saveToFile("my_cards")
	//cardsNew := newDeckFromFile("my")
	fmt.Println("Before")
	cards.print()
	cards.shuffle()
	fmt.Println("After")
	cards.print()

}

func convert(){
	s := "Hi there!"
	b := ([]byte(s))
	fmt.Println(b)
}