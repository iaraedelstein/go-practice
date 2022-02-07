package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck
// which is a slice of string
type deck []string

func newDeck() deck{
	cards := deck{}
	cardsSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardsValues := []string{"Ace", "Two", "Three", "Four"}

	for _,suit := range cardsSuits {
		for _,val := range cardsValues {
			cards = append(cards, val + " of "+ suit)
		}
	}
	return cards;
}

func (d deck) print(){

	for i,card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck,deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() (string){
	return strings.Join([]string(d),",")
}

func (d deck) saveToFile(fileName string) (error){
	data:= []byte(d.toString())
	e := ioutil.WriteFile(fileName, data, 0666)
	return e
}

func newDeckFromFile(fileName string) (deck){
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:",err)
		os.Exit(1)
	}
	str:= string(bs)
	stringSlice := strings.Split(str, ",")
	return deck(stringSlice)
}

func (d deck) shuffle(){
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := range d {
		newPosition := r.Intn(len(d)-1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}