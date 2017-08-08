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

func main() {

	fmt.Println("main.go")

	db, err := sql.Open("postgres", "user=pqgotest dbname=frankcash sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error: Could not establish a connection with the database")
	}

	age := 21

	rows, err := db.Query("SELECT name FROM public.users WHERE age = $1", age)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	user := NewUser()

	for rows.Next() {
		err := rows.Scan(
			&user.Name,
		)
		if err != nil {
			log.Fatal(err)
			break
		}
	}
	fmt.Printf("User's name is %s\n", user.Name)

}
