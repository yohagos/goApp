package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Contact struct (object)
type Contact struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var contacts []Contact

func main() {

	r := mux.NewRouter()

	contacts = append(contacts, Contact{Name: "Yosef", Phone: "06151147741", Email: "yosef@test.de"})
	contacts = append(contacts, Contact{Name: "Michael", Phone: "061512258852", Email: "michael@test.de"})
	contacts = append(contacts, Contact{Name: "Daniel", Phone: "06151369963", Email: "daniel@test.de"})

	r.HandleFunc("/contacts", getContacts).Methods("GET")
	r.HandleFunc("/contacts/{name}", getContact).Methods("GET")
	r.HandleFunc("/contacts", createContact).Methods("POST")
	r.HandleFunc("/contacts/{name}", updateContact).Methods("PUT")
	r.HandleFunc("/contacts/{name}", deleteContact).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":5000", r))
}

func getContacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contacts)
}

func getContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, item := range contacts {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Contact{})
}

func createContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var contact Contact

	_ = json.NewDecoder(r.Body).Decode(&contact)
	contacts = append(contacts, contact)
	json.NewEncoder(w).Encode(contact)
}

func updateContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			var contact Contact
			_ = json.NewDecoder(r.Body).Decode(&contact)
			contact.Name = params["name"]
			contacts = append(contacts, contact)
			json.NewEncoder(w).Encode(contact)
			return
		}
	}
}

func deleteContact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for idx, item := range contacts {
		if item.Name == params["name"] {
			contacts = append(contacts[:idx], contacts[idx+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(contacts)
}
