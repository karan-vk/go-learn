package main

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type SpanishBot struct{}

func main() {
	eb := englishBot{}
	sb := SpanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}
func printGreeting(a bot) {
	println(a.getGreeting())

}

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
