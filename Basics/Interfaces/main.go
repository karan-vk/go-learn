package main

import "fmt"

type englishBot struct{}
type SpanishBot struct{}

func main() {
	eb := englishBot{}
	sb := SpanishBot{}
	printGreeting(eb)
	// printGreeting(sb)

}
func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())

}

// func printGreeting(sb SpanishBot) {
// 	fmt.Println(sb.getGreeting())

// }

func (englishBot) getGreeting() string {
	//
	hello := "hi there"
	return hello

}
func (SpanishBot) getGreeting() string {
	//
	hello := "Hola"
	return hello

}
