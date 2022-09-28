package main

import (
	"encoding/json"
	"flag"
	"html/template"
	"strconv"

	// "io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var plantillas = template.Must(template.ParseGlob("static/*"))

//contactos

type Person struct {
	ID        int       `json:"id,omitempty"`
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

var id = 1

func editPage(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "edit.html", nil)
}

func getPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idGet, _ := strconv.Atoi(params["id"])
	for _, person := range people {
		if person.ID == idGet {
			json.NewEncoder(w).Encode(person)
			break
		}
	}
}

func getPeopleEndPoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	id++
	people = append(people, Person{ID: id, FirstName: r.FormValue("firstname"), LastName: r.FormValue("lastname"), Location: &Location{Country: r.FormValue("country"), City: r.FormValue("city")}, Contact: &Contact{Prefix: r.FormValue("prefix"), Number: r.FormValue("number"), Email: r.FormValue("email")}})
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func deletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idDelete, _ := strconv.Atoi(params["id"])
	for i, person := range people {
		if person.ID == idDelete {
			people = append(people[:i], people[i+1:]...)
			break
		}
	}
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func editPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idEdit, _ := strconv.Atoi(params["id"])
	for i, person := range people {
		if person.ID == idEdit {
			person.FirstName = r.FormValue("firstname")
			person.LastName = r.FormValue("lastname")
			person.Location.Country = r.FormValue("country")
			person.Location.City = r.FormValue("city")
			person.Contact.Prefix = r.FormValue("prefix")
			person.Contact.Number = r.FormValue("number")
			person.Contact.Email = r.FormValue("email")
			people[i] = person
			http.Redirect(w, r, "/", http.StatusMovedPermanently)
			return
		}
	}
}

func main() {
	var dir string

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	router := mux.NewRouter()
	people = append(people, Person{ID: id, FirstName: "Jesus", LastName: "Osorio", Location: &Location{Country: "Colombia", City: "Cali"}, Contact: &Contact{Prefix: "+57", Number: "312312312", Email: "email@gmail.com"}})

	//endpoints
	router.HandleFunc("/people", getPeopleEndPoint)
	router.HandleFunc("/people/edit/{id}", editPage)
	router.HandleFunc("/person/edit/{id}", editPerson)
	router.HandleFunc("/people/crear", createPerson)
	router.HandleFunc("/people/delete/{id}", deletePerson)
	router.HandleFunc("/people/{id}", getPerson)

	router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./static/"))))

	handler := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000", handler))

}
