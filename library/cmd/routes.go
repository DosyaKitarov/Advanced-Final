package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (a *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	router.GET("/books/:id", httprouter.Handle(a.getBookInfo))
	router.POST("/books", httprouter.Handle(a.createBook))
	router.PUT("/books/:id", httprouter.Handle(a.updateBook))
	router.DELETE("/books/:id", httprouter.Handle(a.deleteBook))

	router.GET("/movies/:id", httprouter.Handle(a.getMovie))
	router.POST("/movies", httprouter.Handle(a.createMovie))
	router.PUT("/movies/:id", httprouter.Handle(a.updateMovie))
	router.DELETE("/movies/:id", httprouter.Handle(a.deleteMovie))

	return router
}
