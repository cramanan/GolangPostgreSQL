package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
}

type Storage interface {
	CreateUser(User) error
	ReadUsers() ([]User, error)
	UpdateUser(User, User) error
	DeleteUser(User) error
}

type PostgreSQLStore struct{ *sql.DB }

func (store *PostgreSQLStore) Init() (err error) {
	_, err = store.DB.Exec(
		`CREATE TABLE IF NOT EXISTS users (
		id VARCHAR(36)
	);`)
	return err
}

func NewPostgreSQLStore() (store *PostgreSQLStore, err error) {
	store = new(PostgreSQLStore)
	store.DB, err = sql.Open("postgres",
		fmt.Sprintf(
			"user=%s password=%s sslmode=disable dbname=%s",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_DB"),
		),
	)
	if err != nil {
		return nil, err
	}

	err = store.Init()
	if err != nil {
		return nil, err
	}

	return store, err
}

func (store *PostgreSQLStore) CreateUser(User) (err error) {
	raw, err := uuid.NewV4()
	if err != nil {
		return err
	}
	store.DB.Exec("INSERT INTO users (id) VALUES ($1)", raw.String())
	return
}

func (store *PostgreSQLStore) ReadUsers() (users []User, err error) {
	rows, err := store.DB.Query("SELECT * FROM users;")
	if err != nil {
		return nil, err
	}
	users = []User{}
	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (store *PostgreSQLStore) UpdateUser(user User, update User) (err error) {
	return nil
}

func (store *PostgreSQLStore) DeleteUser(user User) (err error) {
	return nil
}
