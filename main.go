package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Person struct {
	ID        string    `json:"id,omitempty"`
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	Birth     string    `json:"birth,omitempty"`
	Location  *Location `json:"location,omitempty"`
	Contact   *Contact  `json:"contact,omitempty"`
}

type Location struct {
	Country string `json:"country,omitempty"`
	State   string `json:"state,omitempty"`
	City    string `json:"city,omitempty"`
}

type Contact struct {
	Prefix string `json:"prefix,omitempty"`
	Number string `json:"number,omitempty"`
	Email  string `json:"email,omitempty"`
}

var people []Person

func getPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func getPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
			break
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func createPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(person)
}

func deletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for i, person := range people {
		if person.ID == params["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func enviarStatic(w http.ResponseWriter, r *http.Request) {
	file, err := ioutil.ReadFile("./static/index.html")
	if err != nil {
		log.Fatal(err)
	}

	w.Write(file)
}

func main() {
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Jesus", LastName: "Osorio"})

	//endpoints
	router.HandleFunc("/people", getPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", getPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", createPersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", deletePersonEndPoint).Methods("DELETE")

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	router.HandleFunc("/", enviarStatic)

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))

}
