package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitorsName string) string
}

type Italian struct {
}

func (man Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}
func (man Italian) LanguageName() string {
	return "Italian"
}

type German struct {
}

func (man German) Greet(name string) string {
	return fmt.Sprintf("Hallo %s!", name)
}
func (man German) LanguageName() string {
	return "German"
}

type Portuguese struct {
}

func (man Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
func (man Portuguese) LanguageName() string {
	return "Portuguese"
}

func SayHello(name string, holder Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", holder.LanguageName(), holder.Greet(name))
}

func main() {
	fmt.Println(SayHello("Loshara", Italian{}))
}

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
