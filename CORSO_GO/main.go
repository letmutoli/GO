package main

import (
	"CORSO_GO/utenti"
	"fmt"
)

func main() {
	utente := utenti.NuovoUtente("Mario", "mario@mail.com")
	fmt.Println(utente.Nome)
}
