package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"strconv"
	"log"
)

type User struct {
	ID int
	Token string
	Name string
	Age int
}

var Users []User

func userHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	// Проверка является ли id числом 
	id, err := strconv.Atoi(vars["id"])
	if err!= nil {
		log.Println("ERROR parse id %s", err)
		return
	}

	// Проверка существует ли запрашиваемый id в Users
	if id > len(Users) {
		log.Println("ERROR id is not found")
		return
	}

	switch r.Method {
		case http.MethodGet:
			js , err := json.Marshal(Users[id]) 
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(js)
		case http.MethodPost:
			name := r.FormValue("name")
			age, err := strconv.Atoi(r.FormValue("age"))
			if err != nil {
				log.Println("ERROR age is not int")
				return
			}
			// Check validation data
			if name == "" || age < 0 {
				log.Println("ERROR is not validation data")
				return
			}
			Users[id].Name = name
			Users[id].Age = age
	}
}

func main() {

	// Для примера
	Users = append(Users, User{0, "h64eh", "Petya", 19})
	Users = append(Users, User{1, "5h54h", "John", 25})
	Users = append(Users, User{2, "h4yhvf", "Mark", 28})
	Users = append(Users, User{3, "56j33d", "Mike", 35})

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", userHandler)

	http.ListenAndServe(":8080", r)
}