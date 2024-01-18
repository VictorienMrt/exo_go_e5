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
	// Initialize a new dictionary with data stored in 'test.json'.
	d := dictionary.New("test.json")

	// Set up a new router using the Gorilla Mux package.
	r := mux.NewRouter()
	// Apply the LoggerMiddleware to log all incoming requests.
	r.Use(middleware.LoggerMiddleware)

	// Define the '/login' endpoint for user authentication.
	r.HandleFunc("/login", loginHandler).Methods("POST")

	// Create a subrouter for protected routes that require authentication.
	s := r.PathPrefix("/").Subrouter()
	// Apply the AuthMiddleware to all routes under the subrouter.
	s.Use(middleware.AuthMiddleware)
	// Define various endpoints and their corresponding handlers.
	s.HandleFunc("/protected-route", protectedRouteHandler).Methods("GET")
	s.HandleFunc("/entry", addEntryHandler(d)).Methods("POST")
	s.HandleFunc("/entry/{word}", getEntryHandler(d)).Methods("GET")
	s.HandleFunc("/entry/{word}", deleteEntryHandler(d)).Methods("DELETE")

	// Start the server on port 8080 and log any fatal errors.
	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// addEntryHandler handles the addition of new dictionary entries.
func addEntryHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var entry dictionary.Entry
		// Decode the JSON request body into an Entry object.
		err := json.NewDecoder(r.Body).Decode(&entry)
		if err != nil {
			// Respond with an error if the request body is malformed.
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validate the entry data.
		if len(entry.Word) < 3 || len(entry.Definition) < 5 {
			http.Error(w, "Invalid entry: Word and definition must be longer.", http.StatusBadRequest)
			return
		}

		// Add the entry to the dictionary.
		err = d.Add(entry.Word, entry.Definition)
		if err != nil {
			// Respond with an error if adding the entry fails.
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Respond with a status code indicating successful creation.
		w.WriteHeader(http.StatusCreated)
	}
}

// getEntryHandler handles retrieval of a specific dictionary entry.
func getEntryHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the 'word' parameter from the URL.
		vars := mux.Vars(r)
		word := vars["word"]

		// Retrieve the entry from the dictionary.
		entry, err := d.Get(word)
		if err != nil {
			// Respond with an error if the entry is not found.
			http.Error(w, "Entry not found", http.StatusNotFound)
			return
		}
		// Send the entry back in the response body.
		json.NewEncoder(w).Encode(entry)
	}
}

// deleteEntryHandler handles the deletion of a dictionary entry.
func deleteEntryHandler(d *dictionary.Dictionary) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract the 'word' parameter from the URL.
		vars := mux.Vars(r)
		word := vars["word"]

		// Remove the entry from the dictionary.
		err := d.Remove(word)
		if err != nil {
			// Respond with an error if the entry is not found.
			http.Error(w, "Entry not found", http.StatusNotFound)
			return
		}
		// Respond with a status code indicating successful deletion.
		w.WriteHeader(http.StatusOK)
	}
}

// protectedRouteHandler is an example handler for a protected route.
func protectedRouteHandler(w http.ResponseWriter, r *http.Request) {
	// This is a placeholder response for a protected route.
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Access to protected route"))
}
