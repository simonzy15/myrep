package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"myrep/pkg1"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Upvotes   string `json:"upvotes"`
	Downvotes string `json:"downvotes"`
}

type userCreation struct {
	Username string `json:"username"`
}

type userUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type commentStruct struct {
	User_ID     string `json:"user_id"`
	Comment     string `json:"comment"`
	Author_ID   string `json:"author"`
	Author_Name string `json:"author_name"`
}

var DB *sql.DB
var currentUser User

func createUser(w http.ResponseWriter, r *http.Request) {
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

	insertQuery := "INSERT INTO USERS( USER_NAME ) VALUES (?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, insertQuery)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, createdUser.Username)

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
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	fmt.Println(params["username"])

	userinfo, err := DB.Query("select USER_NAME, USER_ID, USER_BIO, USER_UPVOTES, USER_DOWNVOTES from USERS where USER_NAME = ?", params["username"])
	if err != nil {
		// return
		log.Fatal(err)
		return
	}

	defer userinfo.Close()

	for userinfo.Next() {
		err := userinfo.Scan(&currentUser.Username, &currentUser.ID, &currentUser.Bio, &currentUser.Upvotes, &currentUser.Downvotes)
		if err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(currentUser)
		log.Println(currentUser.ID, currentUser.Username)
	}
	return
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

	// if editUser.Username != "" {
	// 	if count > 0 {
	// 		editQuery += ", "
	// 	}
	// 	editQuery += "USER_NAME = \"" + editUser.Username + "\""
	// 	count++
	// }

	editQuery += " WHERE USER_NAME = " + params["username"]

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

	fmt.Println(params["username"])

	if err != nil {
		// return
		log.Fatal(err)
		return
	}
}

func addComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	// params := mux.Vars(r)

	decoder := json.NewDecoder(r.Body)

	var comment commentStruct

	err := decoder.Decode(&comment)

	if err != nil {
		log.Printf("Error %s when preparing DECODING statement", err)
		return
	}

	insertQuery := "INSERT INTO COMMENTS( USER_ID, COMMENTER, COMMENT, COMMENT_TIME) VALUES (?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, insertQuery)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, comment.User_ID, comment.Author_Name, comment.Comment, time.Now())

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

	if err != nil {
		log.Fatal(err)
		return
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	engine := os.Getenv("RDS_ENGINE")
	connectionString := os.Getenv("RDS_CONNECTION_STRING")
	certPath := os.Getenv("CERT_PATH")
	keyPath := os.Getenv("KEY_PATH")

	db, err := sql.Open(engine, connectionString)
	if err != nil {
		log.Fatal("Error loading connecting to SQL")
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
	router.HandleFunc("/api/addcomment", addComment).Methods("POST")
	router.HandleFunc("/api/getuser/{username}", getUser).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/edituser/{username}", editUser).Methods("PUT")

	log.Fatal(http.ListenAndServeTLS(":8001", certPath, keyPath, router))

	pkg1.Add()

	// defer db.Close()
}
