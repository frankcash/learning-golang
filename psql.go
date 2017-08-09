package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID   int
	Name string
	Age  int
	Data json.RawMessage
}

type ShoeDescript struct {
	Favorite string  `json:"favorite"`
	Size     float64 `json:"size"`
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

func getUserByAge(db *sql.DB, age int) *User {
	user := NewUser()

	rows, err := db.Query("SELECT name, data FROM public.users WHERE age = $1", age)
	defer rows.Close()

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		err := rows.Scan(
			&user.Name,
			&user.Data,
		)
		if err != nil {
			log.Fatal(err)
			break
		}
	}

	return user

}

func addUser(db *sql.DB, name string, age int, size float64, favorite string) int64 {
	shoe := ShoeDescript{
		Size:     size,
		Favorite: favorite,
	}
	if b, err := json.Marshal(shoe); err == nil {
		var id int64

		err := db.QueryRow(`INSERT INTO public.users(name,age,data) 
		VALUES ($1,$2,$3) RETURNING id;`, name, age, b).Scan(&id)

		if err != nil {
			fmt.Println("err addUser ", err)
		}

		return id
	}
	return 0

}

func main() {
	fmt.Println("main.go")

	db := connect()
	age := 27
	user := getUserByAge(db, age)
	shoe := ShoeDescript{}
	if err := json.Unmarshal([]byte(user.Data), &shoe); err != nil {
		fmt.Println("Err main ", err)
	}
	fmt.Printf("User's name is %s %v\n", user.Name, shoe)

	resAge := addUser(db, "Vu", 27, 6.5, "ASICS")
	fmt.Printf("Vu is %d\n", resAge)

}
