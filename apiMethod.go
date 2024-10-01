package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func getAllBooksHandler(c *fiber.Ctx) error {
	books, err := getBooks()

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(books)
}

func getBookWithPublisherHandler(c *fiber.Ctx) error {
	bookWithPublisher, err := getBookWithPublisher()

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(bookWithPublisher)
}

func getBookHandler(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	book, err := getBook(bookId)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(book)
}

func createBookHandler(c *fiber.Ctx) error {
	bookCreate := new(Book)

	if err := c.BodyParser(bookCreate); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	book, err := createBook(bookCreate)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(book)

}

func updateBookHandler(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	b := new(Book)

	if err := c.BodyParser(b); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	book, err := updateBook(bookId, b)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(book)

}

func deleteBookHandler(c *fiber.Ctx) error {
	bookId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = deleteBook(bookId)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.SendStatus(fiber.StatusAccepted)

}
