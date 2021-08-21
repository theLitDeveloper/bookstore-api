package bookstore

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// booksHandler is returning all books in store
func booksHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Books)
}

// getBookByIDHandler is returning a single book by ID
func getBookByIDHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return badRequest(ctx)
	}
	for _, book := range Books {
		if book.ID == id {
			return ctx.JSON(http.StatusOK, book)
		}
	}
	return notFound(ctx)
}

// updateBookHandler updates data for a particular book
func updateBookHandler(ctx echo.Context) error {
	var err error

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return badRequest(ctx)
	}

	bdy := ctx.Request().Body
	defer bdy.Close()

	var updBook Book

	err = json.NewDecoder(bdy).Decode(&updBook)
	if err != nil {
		return badRequest(ctx)
	}

	found := false
	for idx, bk := range Books {
		if bk.ID == id {
			found = true
			updBook.ID = id
			Books[idx] = updBook
		}
	}

	if !found {
		return notFound(ctx)
	}

	return ctx.JSON(http.StatusAccepted, updBook)
}

// addBookHandler adds a book to the store
func addBookHandler(ctx echo.Context) error {
	bdy := ctx.Request().Body
	defer bdy.Close()
	var book Book
	err := json.NewDecoder(bdy).Decode(&book)
	if err != nil {
		return badRequest(ctx)
	}
	book.ID = AutoID.ID()
	Books = append(Books, book)

	return ctx.JSON(http.StatusCreated, book)
}

// deleteBookHandler removes a book from the store
func deleteBookHandler(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return badRequest(ctx)
	}

	found := false
	for idx, book := range Books {
		if book.ID == id {
			found = true
			Books = append(Books[:idx], Books[idx+1:]...)
		}
	}
	if !found {
		return notFound(ctx)
	}

	return ctx.NoContent(http.StatusNoContent)
}

// badRequest status 400 HTTP response
func badRequest(ctx echo.Context) error {
	return ctx.String(http.StatusBadRequest, "Bad request")
}

// notFound status 404 HTTP response
func notFound(ctx echo.Context) error {
	resp := make(map[string]string)
	resp["message"] = "Not Found"
	return ctx.JSON(http.StatusNotFound, resp)
}
