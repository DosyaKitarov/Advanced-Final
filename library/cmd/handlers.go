package main

import (
	"encoding/json"
	"fmt"
	grpc "library/pkg/grpc"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type BookResponse struct {
	BookId uint32 `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type MovieResponse struct {
	MovieId  uint32  `json:"movie_id"`
	Title    string  `json:"title"`
	Director string  `json:"director"`
	Rating   float64 `json:"rating"`
}

func (a *application) getBookInfo(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res, err := a.grpc.BookClient.GetBookInfo(r.Context(), &grpc.GetBookInfoRequest{BookId: uint32(id)})
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"book_id":%d,"title":"%s","author":"%s"}`, res.BookId, res.Title, res.Author)
}

func (a *application) createBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var bookData struct {
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	err := json.NewDecoder(r.Body).Decode(&bookData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	res, err := a.grpc.BookClient.CreateBook(r.Context(), &grpc.CreateBookRequest{Title: bookData.Title, Author: bookData.Author})
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(BookResponse{
		BookId: res.BookId,
		Title:  res.Title,
		Author: res.Author,
	})
}

func (a *application) updateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var bookData struct {
		id     uint32
		Title  string `json:"title"`
		Author string `json:"author"`
	}

	err := json.NewDecoder(r.Body).Decode(&bookData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	idStr := ps.ByName("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	bookData.id = uint32(id)

	res, err := a.grpc.BookClient.UpdateBook(r.Context(), &grpc.UpdateBookRequest{BookId: bookData.id, Title: bookData.Title, Author: bookData.Author})
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(BookResponse{
		BookId: res.BookId,
		Title:  res.Title,
		Author: res.Author,
	})
}

func (a *application) deleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = a.grpc.BookClient.DeleteBook(r.Context(), &grpc.DeleteBookRequest{BookId: uint32(id)})
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `Book with ID %d deleted`, id)
}

func (a *application) getMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))

	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	res, err := a.grpc.MovieClient.GetMovie(r.Context(), &grpc.GetMovieRequest{MovieId: uint32(id)})
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `{"movie_id":%d,"title":"%s","director":"%s","rating":%v}`, res.MovieId, res.Title, res.Director, res.Rating)
}

func (a *application) createMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var movieData struct {
		Title    string  `json:"title"`
		Director string  `json:"director"`
		Rating   float64 `json:"rating"`
	}

	err := json.NewDecoder(r.Body).Decode(&movieData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	res, err := a.grpc.MovieClient.CreateMovie(r.Context(), &grpc.CreateMovieRequest{Title: movieData.Title, Director: movieData.Director, Rating: movieData.Rating})
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MovieResponse{
		MovieId:  res.MovieId,
		Title:    res.Title,
		Director: res.Director,
		Rating:   res.Rating,
	})
}

func (a *application) updateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var movieData struct {
		id       uint32
		Title    string  `json:"title"`
		Director string  `json:"director"`
		Rating   float64 `json:"rating"`
	}

	err := json.NewDecoder(r.Body).Decode(&movieData)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	idStr := ps.ByName("id")
	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	movieData.id = uint32(id)

	res, err := a.grpc.MovieClient.UpdateMovie(r.Context(), &grpc.UpdateMovieRequest{MovieId: movieData.id, Title: movieData.Title, Director: movieData.Director, Rating: movieData.Rating})
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(MovieResponse{
		MovieId:  res.MovieId,
		Title:    res.Title,
		Director: res.Director,
		Rating:   res.Rating,
	})
}

func (a *application) deleteMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")

	id, err := strconv.Atoi(strings.TrimSpace(idStr))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	_, err = a.grpc.MovieClient.DeleteMovie(r.Context(), &grpc.DeleteMovieRequest{MovieId: uint32(id)})
	if err != nil {
		a.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, `Movie with ID %d deleted`, id)
}
