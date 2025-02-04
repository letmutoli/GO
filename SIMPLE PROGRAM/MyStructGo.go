package main

import "fmt"

type Persona struct {
	Nome string

	Età int

	Email string
}

func (p Persona) Saluta() string {
	return "Il mio Nome è " + p.Nome
}

func main() {
	p := Persona{
		Nome:  "Lillo",
		Età:   41,
		Email: "let.minutoli@prova.com",
	}

	saluto := p.Saluta()

	fmt.Println(p.Nome)
	fmt.Println(p.Età)
	fmt.Println(p.Email)
	fmt.Println(saluto)

}
