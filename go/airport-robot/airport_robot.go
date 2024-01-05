package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

type Greeter interface {
	LanguageName() string
	Greet(vistorName string) string
}

func SayHello(vistorName string, greeter Greeter) string {
	return greeter.Greet(vistorName)
}

type Italian struct {
}

func (italian Italian) LanguageName() string {
	return "Italian"
}

func (italian Italian) Greet(vistorName string) string {
	return fmt.Sprintf("I can speak %s: Ciao %s!", italian.LanguageName(), vistorName)
}

type Portuguese struct {
}

func (porguese Portuguese) LanguageName() string {
	return "Portuguese"
}

func (porguese Portuguese) Greet(vistorName string) string {
	return fmt.Sprintf("I can speak %s: Olá %s!", porguese.LanguageName(), vistorName)
}
