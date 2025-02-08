package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Title  string
	Header string
	Body   string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{
			Title:  "My Page",
			Header: "Welcome to my page",
			Body:   "This is the body of my page",
		}
		t, err := template.ParseFiles("layout.html", "content.html")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error parsing template files:", err)
			return
		}
		err = t.ExecuteTemplate(w, "content", p)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Println("Error executing template:", err)
		}
	})
	log.Println("Server starting on :8081")
	err := http.ListenAndServe(":8081©", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
