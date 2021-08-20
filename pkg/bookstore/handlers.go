package bookstore

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// BooksHandler is returning a JSON object with all books in store
func BooksHandler(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, Books)
}

func GetBookByIDHandler(ctx echo.Context) error {
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

func UpdateBookHandler(ctx echo.Context) error {
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

func AddBookHandler(ctx echo.Context) error {
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

func DeleteBookHandler(ctx echo.Context) error {
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

func badRequest(ctx echo.Context) error {
	return ctx.String(http.StatusBadRequest, "Bad request")
}

func notFound(ctx echo.Context) error {
	resp := make(map[string]string)
	resp["message"] = "Not Found"
	return ctx.JSON(http.StatusNotFound, resp)
}
