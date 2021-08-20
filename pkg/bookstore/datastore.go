package bookstore

import (
	"sync"
)

var Books []Book
var AutoID autoInc

type autoInc struct {
	sync.Mutex // ensures autoInc is goroutine-safe
	id         int
}

type Book struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Publisher string `json:"publisher"`
	Rating    int    `json:"rating"`
}

// Init adds some books
func initDatastore() {
	Books = append(Books, Book{ID: 1, Title: "Unfuck the Economy", Author: "Waldemar Zeiler", Publisher: "Goldmann Verlag", Rating: 5})
	Books = append(Books, Book{ID: 2, Title: "Reinventing Organizations", Author: "Frederic Laloux", Publisher: "Diateino", Rating: 5})
	Books = append(Books, Book{ID: 3, Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", Publisher: "Pan Books", Rating: 5})
	AutoID.id = 4
}

func (a *autoInc) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}
