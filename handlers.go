package main

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func loginPostHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// username := r.FormValue("")
}

func loginGetHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := store.Get(r, "sid")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	name := mux.Vars(r)["username"]
	if name == "" {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	session.Values["username"] = name
	fmt.Println("session:", session)

	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(session)

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("123")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	user := User{Name: "Sam", Email: "yaob@miamioh.edu", Age: 23, Id: 1}

	jsonResponse, _ := json.Marshal(user)

	w.Write(jsonResponse)

}
