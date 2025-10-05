package main

import (
	"database/sql"
	//_ "github.com/lib/pq" // PostgresSql driver
)

type Store interface {
	GetByID(int) (*User, error)
}
type User struct {
	ID   int
	Name string
}

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id int) (*User, error) {
	row := r.db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)

	var user User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

type InMemRepository struct {
	users []User
}

func (r *InMemRepository) GetByID(id int) (*User, error) {
	return nil, nil
}

//func getUserByID(db *sql.DB, id int) (*User, error) {
//	row := db.QueryRow("SELECT id, name FROM users WHERE id = $1", id)
//
//	var user User
//	err := row.Scan(&user.ID, &user.Name)
//	if err != nil {
//		return nil, err
//	}
//
//	return &user, nil
//}
//
//func main() {
//	connStr := "user=postgres dbname=postgres sslmode=disable"
//	db, err := sql.Open("postgres", connStr)
//	if err != nil {
//		panic(err)
//	}
//	defer db.Close()
//
//	user, err := getUserByID(db, 1)
//	if err != nil {
//		fmt.Println("Error:", err)
//	} else {
//		fmt.Printf("User: %+v/n", user)
//	}
//}
