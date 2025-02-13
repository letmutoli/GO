package main

//import per il progettio
import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

// struttura dati per l'utente per il db
type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

// variabili globali oper il db
var db *sql.DB                        //puntatore a db
var sessionMap = make(map[string]int) //mappa vuota con valore stringa e chiave intera

func init() {
	var err error
	connectionString := "user=postgres password=10Lillo83 dbname=gestione_utenti sslmode=disable"
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connessione al db riuscita")
}

func main() {

	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			t, _ := template.ParseFiles("login.html")
			t.Execute(w, nil)
			return

		}
		email := r.FormValue("email")
		password := r.FormValue("password")

		var user User
		//esecuzione query
		err := db.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1 AND password = $2", email, password).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
		if err != nil {
			http.Error(w, "Credenziali non valide", 400)
		}

		sessionID := uuid.New().String()

		//cookie di sessione
		cookie := http.Cookie{
			Name:  "session_id",
			Value: sessionID,
		}

		http.SetCookie(w, &cookie)
		//add id to session map
		sessionMap[sessionID] = user.Id

		http.Redirect(w, r, "/user", http.StatusSeeOther)
	})

	//percordo register

	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Metodo della richiesta: ", r.Method)
		if r.Method == http.MethodGet {
			t, _ := template.ParseFiles("/login.html")
			t.Execute(w, nil)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		var userID int
		err := db.QueryRow("INSERT INTO users(name, email, password) VALUES($1, $2, $3) RETURNING id", name, email, password).Scan(&userID)

		if err != nil {
			log.Printf("Errore durante la registrazione: %v", err)
			http.Error(w, "Impossibile Registrarsi", 500)
			return
		}

		fmt.Fprintf(w, "Registrazione avvenuta con successo, con ID Utente %d", userID)

	})

	//gestore sessioni
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Non autenticato", http.StatusUnauthorized)
		}
		userID, ok := sessionMap[cookie.Value]

		if !ok {
			http.Error(w, "sessione nin valida", http.StatusUnauthorized)
		}
		var user User
		switch r.Method {
		case http.MethodGet:
			err := db.QueryRow("SELECT id, name, email, password FROM users WHERE id = $1", userID).Scan(&user.Id, &user.Name, &user.Email, &user.Password)
			if err != nil {
				http.Error(w, "Utente non trovato", 404)
				return
			}

			t, _ := template.ParseFiles("/user.html")
			t.Execute(w, user)

		case http.MethodPost:
			// prendo i dati da inserire
			newName := r.FormValue("name")
			newEmail := r.FormValue("email")
			newPassword := r.FormValue("password")

			//aggiorno utente mediante insert
			_, err := db.Exec("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4", newName, newEmail, newPassword, userID)
			if err != nil {
				http.Error(w, "Errore durante l'aggiornamento", 500)
				return
			}
			fmt.Fprint(w, "Utente aggiornato con successo")

		default:
			http.Error(w, "Metodo non supportato", http.StatusMethodNotAllowed)
		}
	})

	//logout
	http.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Non autenticato", http.StatusUnauthorized)
			return
		}
		delete(sessionMap, cookie.Value)
		http.SetCookie(w, &http.Cookie{
			Name:    "session_id",
			Value:   "",
			Expires: time.Unix(0, 0),
		})
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":8080", nil)
}
