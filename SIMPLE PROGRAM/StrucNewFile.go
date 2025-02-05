package main

import "fmt"

type Persona struct {
	Nome string
	Età  int
}

type Impiegato struct {
	Persona
	Posizione string
}

func compleanno(ptr *Persona) {
	ptr.Età++
}

func main() {
	p := Persona{Nome: "Lillo", Età: 40}

	fmt.Println(p.Nome)

	compleanno(&p)

	fmt.Println(p.Età)

	e := Impiegato{Persona: Persona{Nome: "Letterio", Età: 42}, Posizione: "Ingegnere"}

	fmt.Println(e.Nome)
}
