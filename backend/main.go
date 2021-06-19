package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"myrep/pkg1"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
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
var DB *sql.DB
var currentUser User

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

	fmt.Println(params["id"])

	// userinfo, err := db.Query("select USER_ID, USER_NAME from USERS where USER_ID = 12")
	// if err != nil {
	// 	// return
	// 	log.Fatal(err)
	// 	return
	// }

	// defer userinfo.Close()

	// var (
	// 	id   int
	// 	name string
	// )

	// for userinfo.Next() {
	// 	err := userinfo.Scan(&id, &name)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(id, name)
	// }

	// for _, item := range users {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(users)
}

func main() {

	// database connection

	DB, err := sql.Open("mysql", "admin:123234w1nd0w@tcp(myrep-db.cqmo4rbwc6vu.us-east-2.rds.amazonaws.com:3306)/USER_DATA")
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	userinfo, err := DB.Query("select USER_ID, USER_NAME from USERS where USER_ID = 12")
	if err != nil {
		// return
		log.Fatal(err)
		return
	}

	defer userinfo.Close()

	var (
		id   int
		name string
	)

	for userinfo.Next() {
		err := userinfo.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	router := mux.NewRouter()
	router.HandleFunc("/api/register", createUser).Methods("POST")
	router.HandleFunc("/api/getuser/{id}", getUser).Methods("GET")

	log.Fatal(http.ListenAndServe(":8001", router))

	pkg1.Add()

	// defer db.Close()
}
