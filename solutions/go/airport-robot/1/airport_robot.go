package airportrobot

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

func SayHello(name string, greeter Greeter) string {
	return fmt.Sprintf("I can speak %s: %s", greeter.LanguageName(), greeter.Greet(name))
}

type Italian struct {
}

func (s Italian) LanguageName() string {
	return "Italian"
}

func (s Italian) Greet(name string) string {
	return fmt.Sprintf("Ciao %s!", name)
}

type Portuguese struct {
}

func (s Portuguese) LanguageName() string {
	return "Portuguese"
}

func (s Portuguese) Greet(name string) string {
	return fmt.Sprintf("Olá %s!", name)
}
