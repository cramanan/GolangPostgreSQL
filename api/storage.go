package api

import (
	"database/sql"
	"fmt"
	"log"
	"os"

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
	ReadUsers(User) error
	UpdateUser(User, User) error
	DeleteUser(User) error
}

type PostgreSQLStore struct{ *sql.DB }

func (store PostgreSQLStore) Init() (err error) {

	_, err = store.DB.Exec(
		`CREATE TABLE users IF NOT EXISTS (
		id VARCHAR(36)
	);`)
	return err
}

func NewPostgreSQLStore() (store *PostgreSQLStore, err error) {
	store = new(PostgreSQLStore)
	store.DB, err = sql.Open("postgres",
		fmt.Sprintf(
			"user=%s password=%s sslmode=disable",
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_PASSWORD"),
		),
	)
	if err != nil {
		return nil, err
	}

	return store, err
}

func (store *PostgreSQLStore) CreateUser(user User) (err error) {
	return
}

func (store *PostgreSQLStore) ReadUsers(user User) (err error) {
	return nil
}

func (store *PostgreSQLStore) UpdateUser(user User, update User) (err error) {
	return nil
}

func (store *PostgreSQLStore) DeleteUser(user User) (err error) {
	return nil
}
