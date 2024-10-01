package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

var db *sql.DB

type Book struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type BookWithPublisher struct {
	BookID        int
	BookName      string
	BookPrice     int
	PublisherName string
}

func SetupDatabase() *sql.DB {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error!")
	}

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database Connected!")

	return db
}

func main() {
	app := fiber.New()
	db = SetupDatabase()
	defer db.Close()
	app.Get("/books", getAllBooksHandler)
	app.Get("/book-with-publisher", getBookWithPublisherHandler)
	app.Get("/book/:id", getBookHandler)
	app.Post("/book", createBookHandler)
	app.Put("/book/:id", updateBookHandler)
	app.Delete("/book/:id", deleteBookHandler)
	app.Listen(":8080")
}
