package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishChatBot struct{}
type spanishChatBot struct{}

func main() {
	eb := englishChatBot{}
	sb := spanishChatBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishChatBot) getGreeting() string {
	// Very custom logic for getting a greeting
	return "Hello there!"
}

func (spanishChatBot) getGreeting() string {
	return "Hola!"
}
