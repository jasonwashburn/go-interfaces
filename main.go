package main

import "fmt"

type englishChatBot struct{}
type spanishChatBot struct{}

func main() {
	eb := englishChatBot{}
	sb := spanishChatBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(eb englishChatBot) {
	fmt.Println(eb.getGreeting())
}

func printGreeting(sb spanishChatBot) {
	fmt.Println(sp.getGreeting())
}

func (englishChatBot) getGreeting() string {
	// Very custom logic for getting a greeting
	return "Hello there!"
}

func (spanishChatBot) getGreeting() string {
	return "Hola!"
}
