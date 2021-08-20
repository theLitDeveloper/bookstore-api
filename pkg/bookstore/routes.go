package bookstore

import "github.com/labstack/echo/v4"

func initRoutes(e *echo.Echo) {

	e.GET("/books", BooksHandler)
	e.POST("/books", AddBookHandler)
	e.GET("/books/:id", GetBookByIDHandler)
	e.DELETE("/books/:id", DeleteBookHandler)
	e.PATCH("/books/:id", UpdateBookHandler)

}
