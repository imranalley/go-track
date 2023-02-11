package airportrobot

import (
	"fmt"
)

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Italian struct {
	languageName string
	name         string
}

func (i Italian) LanguageName() {
	return fmt.Sprintf("%s", i.languageName)
}

func (i Italian) Greet(name string) {
	return fmt.Sprintf("Ciao %s", name)
}

type Greeter interface {
	LanguageName() string
	Greet(name string) string
}

func SayHello(name string, g Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", g.LanguageName(), g.Greet(name))
}
