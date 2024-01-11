package main

import (
	"encoding/json"
	"exo_go_e5/dictionary"
	"exo_go_e5/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	d := dictionary.New("test.json")
	r := mux.NewRouter()
	r.Use(middleware.LoggerMiddleware)

	r.HandleFunc("/login", loginHandler).Methods("POST")

	s := r.PathPrefix("/").Subrouter()
	s.Use(middleware.AuthMiddleware)
	s.HandleFunc("/protected-route", protectedRouteHandler).Methods("GET")
	s.HandleFunc("/entry", addEntryHandler(d)).Methods("POST")
	s.HandleFunc("/entry/{word}", getEntryHandler(d)).Methods("GET")
	s.HandleFunc("/entry/{word}", deleteEntryHandler(d)).Methods("DELETE")

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func addEntryHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entry dictionary.Entry
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if len(entry.Word) < 3 || len(entry.Definition) < 5 {
			http.Error(w, "Invalid entry: Word and definition must be longer.", http.StatusBadRequest)
			return
		}

		err = d.Add(entry.Word, entry.Definition)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func getEntryHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word := vars["word"]
		entry, err := d.Get(word)
		if err != nil {
			http.Error(w, "Entry not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(entry)
	}
}

func deleteEntryHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		word := vars["word"]
		err := d.Remove(word)
		if err != nil {
			http.Error(w, "Entry not found", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func protectedRouteHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access to protected route"))
}
