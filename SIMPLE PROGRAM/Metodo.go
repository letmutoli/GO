package main

import "fmt"

type Persona struct {
	Nome string
	Età  int
}

func (p Persona) Presenta() {

	fmt.Printf("Ciao sono %s e ho %d anni!!\n", p.Nome, p.Età)
}

func main() {

	let := Persona{"Letterio", 42}
	let.Presenta()

	punto := struct{ x, y int }{x: 1, y: 2}

	fmt.Println(punto.x, punto.y)
}
