package main

import (
	"fmt"
	"net/http"
	"time"
)

func setCookie(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)

	cookie := http.Cookie{
		Name:    "FirstVisit",
		Value:   "1",
		Expires: expiration,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "Cookie settato, vai alla pagina /checkCookie")
}

func checkCookie(w http.ResponseWriter, r *http.Request) {

	cookie, _ := r.Cookie("FirstVisit")
	if cookie != nil {
		fmt.Fprintf(w, "Cookie trovato: %s \n", cookie.Value)
	} else {
		fmt.Fprintf(w, "Cookie non trovato")
	}

}
func main() {
	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/checkCookie", checkCookie)
	http.ListenAndServe(":8080", nil)
}
