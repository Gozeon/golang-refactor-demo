package main

import (
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

var (
	key   = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

func secret(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	fmt.Fprintf(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = true
	session.Save(r, w)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	session.Values["authenticated"] = false
	session.Save(r, w)
}

func main() {
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe("localhost:8080", nil)
}

// curl -s http://localhost:8080/secret
// curl -s -I http://localhost:8080/login
// curl -s --cookie "" http://localhost:8080/secret
