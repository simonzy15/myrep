package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
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
	Picture   string `json:"picture"`
}

type userCreation struct {
	Username string `json:"username"`
	Picture  string `json:"picture"`
}

type userUpdate struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type userVote struct {
	User_ID   string `json:"user_id"`
	Author_ID string `json:"author"`
	Vote      int    `json:"vote"` // 0 for downvote, 1 for upvote
}

type commentData struct {
	TargetUser string `json:"target"`
	Commenter  string `json:"commenter"`
	Comment    string `json:"comment"`
	Time       string `json:"time"`
}

var DB *sql.DB
var currentUser User
var comment commentData

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

	insertQuery := "INSERT INTO USERS( USER_NAME, USER_PICTURE ) VALUES (?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, insertQuery)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, createdUser.Username, createdUser.Picture)

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

	userinfo, err := DB.Query("select USER_NAME, USER_ID, USER_BIO, USER_UPVOTES, USER_DOWNVOTES, USER_PICTURE from USERS where USER_NAME = ?", params["username"])
	if err != nil {
		// return
		log.Fatal(err)
		return
	}

	defer userinfo.Close()

	for userinfo.Next() {
		err := userinfo.Scan(&currentUser.Username, &currentUser.ID, &currentUser.Bio, &currentUser.Upvotes, &currentUser.Downvotes, &currentUser.Picture)
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
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, PUT")

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

	editQuery += " WHERE USER_NAME = \"" + params["username"] + "\""

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

	decoder := json.NewDecoder(r.Body)

	var comment commentData

	err := decoder.Decode(&comment)

	if err != nil {
		log.Printf("Error %s when preparing DECODING statement", err)
		return
	}

	insertQuery := "INSERT INTO COMMENTS( USER_NAME, COMMENTER, COMMENT, COMMENT_TIME) VALUES (?, ?, ?, ?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, insertQuery)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, comment.TargetUser, comment.Commenter, comment.Comment, time.Now())

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

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	commentInfo, err := DB.Query("select USER_NAME, COMMENTER, COMMENT, COMMENT_TIME from COMMENTS where USER_NAME = ?", params["target"])
	if err != nil {
		log.Fatal(err)
		return
	}

	defer commentInfo.Close()
	comments := make([]commentData, 0)

	for commentInfo.Next() {
		err := commentInfo.Scan(&comment.TargetUser, &comment.Commenter, &comment.Comment, &comment.Time)
		if err != nil {
			log.Fatal(err)
		}
		comments = append(comments, comment)
	}
	json.NewEncoder(w).Encode(comments)
	return
}

func addVote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")

	decoder := json.NewDecoder(r.Body)

	var vote userVote

	err := decoder.Decode(&vote)

	if err != nil {
		log.Printf("Error %s when preparing DECODING statement", err)
		return
	}

	insertQuery := "INSERT INTO VOTES( VOTE_KEY, VOTE_AUTHOR, VOTE_USER, VOTE) VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE VOTE=(?)"

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancelfunc()

	stmt, err := DB.PrepareContext(ctx, insertQuery)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return
	}

	defer stmt.Close()

	res, err := stmt.ExecContext(ctx, vote.Author_ID+vote.User_ID, vote.Author_ID, vote.User_ID, vote.Vote, vote.Vote)

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

	// Now we need to update the counts in the  user table by calling the stored procedure

	sp_call := "CALL UPDATE_VOTE((?))"

	stmt, err = DB.PrepareContext(ctx, sp_call)

	if err != nil {
		log.Printf("Error %s when preparing SP_SQL statement", err)
		return
	}

	res, err = stmt.ExecContext(ctx, vote.User_ID)

	if err != nil {
		log.Printf("Error %s when executing SP_SQL statement", err)
		return
	}

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
	router.HandleFunc("/api/getcomments/{target}", getComments).Methods("GET")
	router.HandleFunc("/api/addvote", addVote).Methods("POST")
	router.HandleFunc("/api/getuser/{username}", getUser).Methods("GET")
	router.HandleFunc("/api/edituser/{username}", editUser).Methods("PUT", "OPTIONS")

	log.Fatal(http.ListenAndServeTLS(":8001", certPath, keyPath, router))

	// defer db.Close()
}
