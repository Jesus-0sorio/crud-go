package main

import (
	"encoding/json"
	"flag"
	"html/template"

	// "io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var plantillas = template.Must(template.ParseGlob("static/*"))

type Person struct {
	ID        string    `json:"id,omitempty"`
	FirstName string    `json:"firstname,omitempty"`
	LastName  string    `json:"lastname,omitempty"`
	Location  *Location `json:"location,omitempty"`
	Contact   *Contact  `json:"contact,omitempty"`
}

type Location struct {
	Country string `json:"country,omitempty"`
	City    string `json:"city,omitempty"`
}

type Contact struct {
	Prefix string `json:"prefix,omitempty"`
	Number string `json:"number,omitempty"`
	Email  string `json:"email,omitempty"`
}

var people []Person

func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, person := range people {
		if person.ID == params["id"] {
			json.NewEncoder(w).Encode(person)
			break
		}
	}
}

func getPeopleEndPoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func createPersonEndPoint(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	var res string
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	for _, person := range people {
		if person.ID == id {
			res = "SON IGUALES"
			break
		}
	}
	people = append(people, Person{ID: id, FirstName: r.FormValue("firstname"), LastName: r.FormValue("lastname"), Location: &Location{Country: r.FormValue("country"), City: r.FormValue("city")}, Contact: &Contact{Prefix: r.FormValue("prefix"), Number: r.FormValue("number"), Email: r.FormValue("email")}})
	if res == "SON IGUALES" {
		http.RedirectHandler("/error", http.StatusMovedPermanently)
		people = append(people[:len(people)-1], people[len(people):]...)
	}
	http.RedirectHandler("/", http.StatusMovedPermanently)
}

func deletePersonEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, person := range people {
		if person.ID == params["id"] {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func editPerson(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	params := mux.Vars(r)
	log.Println(id)
	for _, person := range people {
		if person.ID == params["id"] {
			log.Println(r.FormValue("firstname"), r.FormValue("lastname"))
			person.FirstName= r.FormValue("firstname")
			person.LastName =  r.FormValue("lastname")
			person.Location.Country = r.FormValue("country")
			person.Location.City = r.FormValue("city")
			person.Contact.Prefix = r.FormValue("prefix")
			person.Contact.Number = r.FormValue("number")
			person.Contact.Email = r.FormValue("email")
			return
		}
	}
}

func editPage(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "edit.html", nil)
}

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	router := mux.NewRouter()

	people = append(people, Person{ID: "1", FirstName: "Jesus", LastName: "Osorio", Location: &Location{Country: "Colombia", City: "Cali"}, Contact: &Contact{Prefix: "+57", Number: "1323468", Email: "jesus@gmail.com"}})
	people = append(people, Person{ID: "2", FirstName: "Andres", LastName: "Osorio", Location: &Location{Country: "Colombia", City: "Cali"}, Contact: &Contact{Prefix: "+57", Number: "1323468", Email: "jesus@gmail.com"}})
	people = append(people, Person{ID: "3", FirstName: "Andres", LastName: "Osorio", Location: &Location{Country: "Colombia", City: "Cali"}, Contact: &Contact{Prefix: "+57", Number: "1323468", Email: "jesus@gmail.com"}})
	people = append(people, Person{ID: "4", FirstName: "Andres", LastName: "Osorio", Location: &Location{Country: "Colombia", City: "Cali"}, Contact: &Contact{Prefix: "+57", Number: "1323468", Email: "jesus@gmail.com"}})
	people = append(people, Person{ID: "5", FirstName: "Andres", LastName: "Osorio", Location: &Location{Country: "Colombia", City: "Cali"}, Contact: &Contact{Prefix: "+57", Number: "1323468", Email: "jesus@gmail.com"}})
	people = append(people, Person{ID: "6", FirstName: "Andres", LastName: "Osorio", Location: &Location{Country: "Colombia", City: "Cali"}, Contact: &Contact{Prefix: "+57", Number: "1323468", Email: "jesus@gmail.com"}})
	people = append(people, Person{ID: "7", FirstName: "Andres", LastName: "Osorio", Location: &Location{Country: "Colombia", City: "Cali"}, Contact: &Contact{Prefix: "+57", Number: "1323468", Email: "jesus@gmail.com"}})

	//endpoints
	router.HandleFunc("/people", getPeopleEndPoint)
	router.HandleFunc("/people/edit/{id}", editPage)
	router.HandleFunc("/person/edit/{id}", editPerson)
	router.HandleFunc("/people/crear", createPersonEndPoint)
	router.HandleFunc("/people/delete/{id}", deletePersonEndPoint)
	router.HandleFunc("/people/{id}", getPerson)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))
	router.PathPrefix("/people/edit/{id}").Handler(http.StripPrefix("/people/edit/{id}", http.FileServer(http.Dir("./static/"))))

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))

}
