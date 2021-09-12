package main

import "fmt"

/**
@bot interface
if you are a type in this program with a function called getGreeting() and you return a string,
then you are a honorary member of type 'bot'
*/
type bot interface {
	getGreeting() string
}

type englishBot struct{} // this is also of type `bot`
type spanishBot struct{} // this is also of type `bot`

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

/**
@printGreeting() - function which accepts anything of type bot interface.
instead of the above two functions, we wrote a single fucntion with the use of an interface.
*/
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	return "Hello!"
}

func (sb spanishBot) getGreeting() string {
	return "Halo!"
}
