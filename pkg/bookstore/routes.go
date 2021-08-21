package bookstore

import "github.com/labstack/echo/v4"

func initRoutes(e *echo.Echo) {

	e.GET("/books", booksHandler)
	e.POST("/books", addBookHandler)
	e.GET("/books/:id", getBookByIDHandler)
	e.DELETE("/books/:id", deleteBookHandler)
	e.PATCH("/books/:id", updateBookHandler)

}
