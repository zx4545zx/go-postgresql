package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	ID int
	Username string
	Password string
}

type Users []User

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "zx4545zx"
	dbname   = "Testdb"
)

func dbConnect() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "HOME")
}

func ShowUsers(w http.ResponseWriter, r *http.Request) {
	user := Users{User{
		ID: 1,
		Username: "hello",
		Password: "zx4545zx",
	}}

	json.NewEncoder(w).Encode(user)
}

func AddUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test POST")
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/user", ShowUsers).Methods("GET")
	myRouter.HandleFunc("/user", AddUsers).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {
	dbConnect()
	handleRequest()
}
