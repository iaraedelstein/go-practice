package main

import "fmt"

//structs with no fields porpuse of create functions as receivers to this structs
type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main () {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//no es posibile sobrecargar los metodos como en java
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }
// func printGreeting(sb spanishhBot) {
// 	fmt.Println(sb.getGreeting())
// }

//use of interface
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//VERY custom logic for generating an english greeting
	return "Hi there!"
}

//podemos omitir inicializar la variable de (sb spanishBot) porque no la estaríamos usando
func (spanishBot) getGreeting() string {
	//VERY custom logic for generating an spanish greeting
	return "¡Hola!"
}