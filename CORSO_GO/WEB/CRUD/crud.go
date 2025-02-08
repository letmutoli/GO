package main

import (
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
}

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&User{})
}

func main() {

	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		db.Create(&User{Name: name})
		fmt.Fprintf(w, "User created")
	})

	http.HandleFunc("/read", func(w http.ResponseWriter, r *http.Request) {
		var users []User
		db.Find(&users)
		fmt.Fprintf(w, "Lista Utenti:\n")
		for _, user := range users {
			fmt.Fprintf(w, user.Name, "\n")
		}
	})

	http.Handle("/update", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		name := r.URL.Query().Get("name")

		var user User
		db.First(&user, id)
		user.Name = name
		db.Save(&user)
		fmt.Fprintf(w, "User updated")
	}))

	http.Handle("/delete", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		var user User
		db.First(&user, id)
		db.Delete(&user)
		fmt.Fprintf(w, "User deleted")
	}))
	http.ListenAndServe(":8080", nil)
}
