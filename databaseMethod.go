package main

func createBook(book *Book) (Book, error) {
	var createdBook Book
	row := db.QueryRow(
		"INSERT INTO public.books(name, price) VALUES ($1, $2) RETURNING id, name, price;;",
		book.Name,
		book.Price,
	)

	err := row.Scan(&createdBook.ID, &createdBook.Name, &createdBook.Price)

	if err != nil {
		return Book{}, err
	}

	return createdBook, nil
}

func getBook(id int) (Book, error) {
	var b Book
	row := db.QueryRow(
		"SELECT id,name,price FROM books WHERE id=$1;",
		id,
	)

	err := row.Scan(&b.ID, &b.Name, &b.Price)

	if err != nil {
		return Book{}, err
	}

	return b, nil

}

func getBooks() ([]Book, error) {
	rows, err := db.Query("SELECT id, name, price FROM books")
	if err != nil {
		return nil, err
	}

	var books []Book

	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Name, &b.Price)

		if err != nil {
			return nil, err
		}

		books = append(books, b)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func getBookWithPublisher() ([]BookWithPublisher, error) {
	query := `
		SELECT
			books.id,
			books.name,
			books.price,
			publishers.name
		FROM
			public.books
		INNER JOIN publishers
			ON books.publisher_id = publishers.id;`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []BookWithPublisher
	for rows.Next() {
		var bookWithPublisher BookWithPublisher

		err := rows.Scan(&bookWithPublisher.BookID, &bookWithPublisher.BookName, &bookWithPublisher.BookPrice, &bookWithPublisher.PublisherName)
		if err != nil {
			return nil, err
		}
		books = append(books, bookWithPublisher)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func updateBook(id int, book *Book) (Book, error) {
	var updatedBook Book
	row := db.QueryRow(
		"UPDATE public.books SET name=$1, price=$2 WHERE id=$3 RETURNING id, name ,price;",
		book.Name,
		book.Price,
		id,
	)

	err := row.Scan(&updatedBook.ID, &updatedBook.Name, &updatedBook.Price)

	if err != nil {
		return Book{}, err
	}

	return updatedBook, nil
}

func deleteBook(id int) error {
	_, err := db.Exec(
		"DELETE FROM public.books WHERE id=$1;",
		id,
	)
	return err
}
