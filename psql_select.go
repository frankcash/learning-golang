package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func NewUser() *User {
	u := User{}
	return &u
}

func connect() *sql.DB {
	db, err := sql.Open("postgres", "user=pqgotest dbname=frankcash sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}
	return db
}

func getUser(db *sql.DB, age int) *User {
	user := NewUser()

	rows, err := db.Query("SELECT name FROM public.users WHERE age = $1", age)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(
			&user.Name,
		)
		if err != nil {
			log.Fatal(err)
			break
		}
	}

	return user

}

func main() {
	fmt.Println("main.go")

	db := connect()
	age := 21
	user := getUser(db, age)
	fmt.Printf("User's name is %s\n", user.Name)

}
