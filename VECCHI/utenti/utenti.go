package utenti

type Utente struct {
	Nome  string
	Email string
}

func NuovoUtente(nome string, email string) *Utente {
	return &Utente{
		Nome:  nome,
		Email: email,
	}
}
