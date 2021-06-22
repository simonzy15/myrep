package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"myrep/pkg1"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

// type Register struct {
// 	User  *User  `jason:"user"`
// 	Email string `json:"email"`
// }

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type userCreation struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type userUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

// var users []User
var DB *sql.DB
var currentUser User

func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	decoder := json.NewDecoder(r.Body)

	var createdUser userCreation

	err := decoder.Decode(&createdUser)

	if err != nil {
		log.Printf("Error %s when preparing Decoding statement", err)
		return
	}

	insertQuery := "INSERT INTO USERS( USER_ID, USER_NAME, USER_BIO, USER_EMAIL ) VALUES (?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, insertQuery)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, createdUser.ID, createdUser.Username, createdUser.Bio, createdUser.Email)

	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return
	}
	log.Printf("%d products created ", rows)

	// var user User
	// _ = json.NewDecoder(r.Body).Decode(&user)
	// user.ID = strconv.Itoa(rand.Intn(100000000))
	// users = append(users, user)
	// json.NewEncoder(w).Encode(user)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// json.NewEncoder(w).Encode(users)

	fmt.Println(params["id"])

	userinfo, err := DB.Query("select USER_NAME, USER_ID from COMMENTS where USER_ID = ?", params["id"])
	if err != nil {
		// return
		log.Fatal(err)
		return
	}

	defer userinfo.Close()

	for userinfo.Next() {
		err := userinfo.Scan(&currentUser.ID, &currentUser.Username)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(currentUser.ID, currentUser.Username)
	}

	// for _, item := range users {
	// 	if item.ID == params["id"] {
	// 		json.NewEncoder(w).Encode(item)
	// 		return
	// 	}
	// }
	// json.NewEncoder(w).Encode(users)
}

func editUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	decoder := json.NewDecoder(r.Body)

	var editUser userUpdate

	count := 0

	err := decoder.Decode(&editUser)

	// USER_NAME = ?, USER_BIO = ?, USER_EMAIL = ? where USER_ID = ?

	editQuery := "UPDATE USERS SET "

	if editUser.Email != "" {
		editQuery += "USER_EMAIL = \"" + editUser.Email + "\""
		count++
	}

	if editUser.Bio != "" {
		if count > 0 {
			editQuery += ", "
		}
		editQuery += "USER_BIO = \"" + editUser.Bio + "\""
		count++
	}

	if editUser.Username != "" {
		if count > 0 {
			editQuery += ", "
		}
		editQuery += "USER_NAME = \"" + editUser.Username + "\""
		count++
	}

	editQuery += " WHERE USER_ID = " + params["id"]

	fmt.Println(editQuery)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, editQuery)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx)

	if err != nil {
		log.Printf("Error %s when executing SQL statement", err)
		return
	}

	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return
	}
	log.Printf("%d products edited ", rows)

	fmt.Println(params["id"])

	if err != nil {
		// return
		log.Fatal(err)
		return
	}

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

	db, err := sql.Open("mysql", "admin:123234w1nd0w@tcp(myrep-db.cqmo4rbwc6vu.us-east-2.rds.amazonaws.com:3306)/USER_DATA")
	if err != nil {
		panic(err)
	}

	DB = db

	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// See "Important settings" section.
	DB.SetConnMaxLifetime(time.Minute * 3)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)

	router := mux.NewRouter()
	router.HandleFunc("/api/register", createUser).Methods("POST")
	router.HandleFunc("/api/getuser/{id}", getUser).Methods("GET")
	router.HandleFunc("/api/edituser/{id}", editUser).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8001", router))

	pkg1.Add()

	// defer db.Close()
}
