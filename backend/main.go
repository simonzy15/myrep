package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Register struct {
	User  *User  `jason:"user"`
	Email string `json:"email"`
}

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users []User

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.ID = strconv.Itoa(rand.Intn(100000000))
	users = append(users, user)
	json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	json.NewEncoder(w).Encode(users)
	for _, item := range users {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/register", createUser).Methods("POST")
	router.HandleFunc("/api/getuser/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", router))
}
