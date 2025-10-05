package main

import (
	"database/sql"
	"fmt"
)

type application struct {
	store Store
}

type UserRepository interface {
	GetById(id int) (*User, error)
}

func main() {
	connStr := "user=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	userRepository := NewPostgresUserRepository(db)
	userService := NewUserService(userRepository)

	inMemRepository := &InMemRepository{}

	app := &application{
		store: userRepository,
		//store: inMemRepository,
	}

	user, err := userService.GetUserByID(1)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("User: %+v/n", user)
	}

}
